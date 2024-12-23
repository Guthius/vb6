For i = 1 To MAX_PLAYERS
    If IsConnected(i) And Trim(GetPlayerIP(i)) = Trim(IP) Then
        n = n + 1
        
        If (n > 1) Then
            IsMultiIPOnline = True
            Exit Function
        End If
    End If
Next i