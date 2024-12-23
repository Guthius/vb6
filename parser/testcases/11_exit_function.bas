Function IsMultiAccounts(ByVal Login As String) As Boolean
Dim i As Long

    If IsConnected(i) And LCase(Trim(Player(i).Login)) = LCase(Trim(Login)) Then
        IsMultiAccounts = True
        Exit Function
    End If
End Function