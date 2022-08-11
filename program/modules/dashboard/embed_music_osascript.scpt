if application "{{PLAYER}}" is running then
    tell application "{{PLAYER}}"
        if player state is playing then
            return name of current track & " ः " & artist of current track & " ः " & album of current track & " ः " & duration of current track & " ः " & {player position}  & " ः " 
        end if
    end tell
end if