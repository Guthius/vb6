If MapNpc(y, x).y > GetPlayerY(Target) And DidWalk = False Then
    If CanNpcMove(y, x, DIR_UP) Then
        Call NpcMove(y, x, DIR_UP, MOVING_WALKING)
        DidWalk = True
    ElseIf CanNpcMove(y, x, DIR_LEFT) Then
        Call NpcMove(y, x, DIR_LEFT, MOVING_WALKING)
        DidWalk = True
    End If
ElseIf MapNpc(y, x).y < GetPlayerY(Target) And DidWalk = False Then
    If CanNpcMove(y, x, DIR_DOWN) Then
        Call NpcMove(y, x, DIR_DOWN, MOVING_WALKING)
        DidWalk = True
    ElseIf CanNpcMove(y, x, DIR_RIGHT) Then
        Call NpcMove(y, x, DIR_RIGHT, MOVING_WALKING)
        DidWalk = True
    End If
ElseIf MapNpc(y, x).x > GetPlayerX(Target) And DidWalk = False Then
    If CanNpcMove(y, x, DIR_LEFT) Then
        Call NpcMove(y, x, DIR_LEFT, MOVING_WALKING)
        DidWalk = True
    ElseIf CanNpcMove(y, x, DIR_UP) Then
        Call NpcMove(y, x, DIR_UP, MOVING_WALKING)
        DidWalk = True
    End If
Else
    Print("Hello World")
End If