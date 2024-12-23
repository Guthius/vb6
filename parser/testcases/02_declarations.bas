Public SpawnSeconds As Long
Public GameName As String * 20 = "Hello World"

Public Const MAX_SPAWN_SECONDS = 10

Const MAX_MAPS = 2  * 20 + MAX_SPAWN_SECONDS

Public Const START_MAP = 1
Private Const START_X = MAX_MAPX / 2
Public Const START_Y = MAX_MAPY / 2

Public Const ADMIN_LOG = "admin.txt"
Public Const PLAYER_LOG = "player.txt"

Public Class() As ClassRec
Public Item(1 To MAX_ITEMS) As ItemRec
Public Npc(1 To MAX_NPCS) As NpcRec
