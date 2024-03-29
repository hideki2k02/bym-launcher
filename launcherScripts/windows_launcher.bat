@echo off

REM Define API endpoint
set API_ENDPOINT=https://api.bymrefitted.com/launcher.json
set "JSON_FILE=./launcher.json"
set "SWF_TO_LAUNCH=stable"

REM Fetch JSON response from API endpoint
powershell -Command "(Invoke-WebRequest -Uri '%API_ENDPOINT%').Content" > launcher.json

set "BUILD_FOLDER=.\builds"
set "FLASH_RUNTIME_FOLDER=.\runtimes"

:: Check if the JSON file exists
if not exist "%JSON_FILE%" (
    echo "JSON file not found: %JSON_FILE%"
    exit /b 1
)

:: Use PowerShell to parse JSON and get the value for "currentLauncherVersion"
for /f "usebackq tokens=*" %%I in (`powershell -Command "(Get-Content '%JSON_FILE%' | ConvertFrom-Json).currentLauncherVersion"`) do (
    @REM echo Target Launcher version: %%I
    set currentLauncherVersion=%%I
)

REM Compare currentLauncherVersion with expected version
set expectedVersion=0.1.0

if "%currentLauncherVersion%" neq "%expectedVersion%" (
    echo "Error: Current launcher version (%currentLauncherVersion%) does not match expected version (%expectedVersion%), please download the latest launcher version"
    exit /b 1
) else (
    echo Launcher version is up to date.
)

:: Create the build folder if it doesn't exist
if not exist "%BUILD_FOLDER%" mkdir "%BUILD_FOLDER%"

if not exist "%FLASH_RUNTIME_FOLDER%" mkdir "%FLASH_RUNTIME_FOLDER%"

:: Function to download missing SWF files
:downloadMissingSWFs
@REM     @REM TODO: Clean up old swf versions
for /f "tokens=1,2 delims=," %%a in ('powershell -Command "(Get-Content '%JSON_FILE%' | ConvertFrom-Json).builds.PSObject.properties | ForEach-Object { $_.Name + ',' + $_.Value }"') do (
    @REM echo %%a ves
    @REM echo Value: %%b
    if "%%a"=="%SWF_TO_LAUNCH%" (
        set "SWF_FILE_TO_LAUNCH=%%b"
    )
    if not exist "%BUILD_FOLDER%\%%b" (
        echo Downloading missing SWF: %BUILD_FOLDER%\%%b
        powershell -Command "(New-Object System.Net.WebClient).DownloadFile('https://api.bymrefitted.com/launcher/downloads/%%b', '%BUILD_FOLDER%\%%b')"
    )
)
goto :downloadFlashRuntime

:downloadFlashRuntime
for /f "tokens=*" %%c in ('powershell -Command "$json = Get-Content '%JSON_FILE%' -Raw | ConvertFrom-Json;if($json.flashRuntimes.windows){Write-Output $json.flashRuntimes.windows}else{Write-Error 'Runtime not found'; exit 1}"') do (
    if not exist "%FLASH_RUNTIME_FOLDER%\%%c" (
        echo Downloading flash runtime for Windows: %%c
        powershell -Command "(New-Object System.Net.WebClient).DownloadFile('http://api.bymrefitted.com/launcher/downloads/%%c', '%FLASH_RUNTIME_FOLDER%\%%c')"
    )
    set LAUNCH_RUNTIME=%FLASH_RUNTIME_FOLDER%\%%c
)
goto :runSWF

:runSWF
REM After all operations, launch the stable SWF using the downloaded Flash runtime executable
%LAUNCH_RUNTIME% "%CD%\builds\%SWF_FILE_TO_LAUNCH%"
goto :eof



:: Call function to download missing SWF files
call :downloadMissingSWFs

goto :eof