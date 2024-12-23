' Check for weapon
If GetPlayerWeaponSlot(Attacker) > 0 Then
    n = GetPlayerInvItemNum(Attacker, GetPlayerWeaponSlot(Attacker))
Else
    n = 0
End If