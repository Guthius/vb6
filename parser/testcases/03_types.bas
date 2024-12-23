Type PlayerRec
    ' General
    Name As String * NAME_LENGTH
    Class As Byte
    Sprite As Integer
    Level As Byte
    Exp As Long
    Access As Byte
    PK As Byte
    
    ' Vitals
    HP As Long
    MP As Long
    SP As Long
    
    ' Stats
    STR As Byte
    DEF As Byte
    SPEED As Byte
    MAGI As Byte
    POINTS As Byte
    
    ' Worn equipment
    ArmorSlot As Byte
    WeaponSlot As Byte
    HelmetSlot As Byte
    ShieldSlot As Byte
    
    ' Inventory
    Inv(1 To MAX_INV) As PlayerInvRec
    Spell(1 To MAX_PLAYER_SPELLS) As Byte
       
    ' Position
    Map As Integer
    x As Byte
    y As Byte
    Dir As Byte
    
    ' Client use only
    MaxHP As Long
    MaxMP As Long
    MaxSP As Long
    XOffset As Integer
    YOffset As Integer
    Moving As Byte
    Attacking As Byte
    AttackTimer As Long
    MapGetTimer As Long
    CastedSpell As Byte
End Type