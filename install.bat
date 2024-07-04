@echo off
mkdir C:\YoutubeMusic
move YoutubeMusic.exe C:\YoutubeMusic
move ym.ico C:\YoutubeMusic 

set SCRIPT="%TEMP%\%RANDOM%-%RANDOM%-%RANDOM%-%RANDOM%.vbs"

echo Set oWS = WScript.CreateObject("WScript.Shell") >> %SCRIPT%
echo sLinkFile = "%USERPROFILE%\Desktop\YoutubeMusic.lnk" >> %SCRIPT%
echo Set oLink = oWS.CreateShortcut(sLinkFile) >> %SCRIPT%
echo oLink.TargetPath = "C:\YoutubeMusic\YoutubeMusic.exe" >> %SCRIPT%
echo oLink.IconLocation = "C:\YoutubeMusic\ym.ico" >> %SCRIPT%
echo oLink.WorkingDirectory = "C:\YoutubeMusic" >> %SCRIPT%
echo oLink.Save >> %SCRIPT%

cscript /nologo %SCRIPT%
del %SCRIPT%
