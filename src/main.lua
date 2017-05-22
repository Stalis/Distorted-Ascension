function extend(child, parent)
  setmetatable(child, {__index = parent})
end

TileSource = {'#', '@', '%'}
sTileTypes = { "Land1", "Land2", "Land3"}
eTileTypes = {}

  
cTileType = {}
function cTileType:new(Type, Prevalence, Source)
  --Type Checking
  if type(Type) ~= "string" then return nil end
  if #Type == 0 then return nil end
  if type(Prevalence) ~= "number" then return nil end
  if type(Source) ~= "number" then return nil end
  if Source > #TileSource then return nil end
  
  local public = {}
  public.Name = Type
  public.Prevalence = Prevalence
  public.Source = TileSource[Source]
  
  setmetatable(public, self)
  self.__index = self
  return public
end

cTile = {}
function cTile:new(TileType)
  --Type Checking
  if type(TileType) ~= "string" then return nil end
  
  local typecheck = false
  local index
  for i = 1, #eTileTypes do
    if TileType == eTileTypes[i].Name then typecheck = true; index = i; break end
  end
  if not typecheck then return nil else typecheck = nil end
  
  local private = {}
    private.Type = eTileTypes[index]
  local public = {}
    function public:GetSource() return private.Type.Source end
    function public:GetType() return private.Type.Name end
    function public:GetPrevalence() return private.Type.Prevalence end
    
  setmetatable(public, self)
  self.__index = self
  return public
end

cChunk = {}
function cChunk:new()
  local private = {}
    private.Map = {}
    
  local public = {}
    --Getters
    function public:GetMap() return private.Map end
    
    
    function public:Generate()
      math.randomseed(os.time())
      local WeightSum = 0
      for i = 1, #eTileTypes do
        WeightSum = WeightSum + eTileTypes[i].Prevalence
      end
    
      for i = 1, 32 do
        local row = {}
        for j = 1, 32 do
          local r = math.random(1, WeightSum) 
          local CurrWeight = 0
          
          for k = 1, #sTileTypes do
            CurrWeight = CurrWeight + eTileTypes[k].Prevalence
            if CurrWeight >= r then
              row[j] = cTile:new(eTileTypes[k].Name)
              break
            end
          end
        end
        private.Map[i] = row

      end
      return true
    end
    
    function public:PrintMap()
      if #private.Map == 0 then return nil end
      for _, row in pairs(private.Map) do
        local outRow = ""
        for _, tile in pairs(row) do
          outRow = outRow..tile:GetSource()..' '
        end
        print(outRow)
      end
      return true
    end
    

  
  setmetatable(public, self)
  self.__index = self
  return public
end



eTileTypes[1] = cTileType:new(sTileTypes[1], 50, 1)
eTileTypes[2] = cTileType:new(sTileTypes[2], 130, 2)
eTileTypes[3] = cTileType:new(sTileTypes[3], 20, 3)

local Chunk = cChunk:new()
Chunk:Generate()

--#################################
-- Code, related to love engine
function DrawMap(Chunk)
  love.graphics.setBackgroundColor(0,0,0)
  local Map = Chunk:GetMap()
  if #Map == 0 then return nil end
  local coords = {x = 1, y = 1}
  for y, row in pairs(Map) do
    for x, tile in pairs(row) do
      v = tile:GetSource()
      if     v == '#' then love.graphics.setColor(144, 173, 0)
      elseif v == '@' then love.graphics.setColor(78, 88, 155)
      elseif v == '%' then love.graphics.setColor(209, 178, 200)
      end
      love.graphics.rectangle('fill', coords.x, coords.y, 16, 16)
      coords.x = coords.x + 17
    end
    coords.x = 1
    coords.y = coords.y + 17
  end
  return true
end

function love.draw() pcall(DrawMap(Chunk)) end 