@echo off
setlocal enabledelayedexpansion

:: Definir las carpetas
set "folders=cmd config internal\database internal\handlers internal\services pkg scripts"

:: Definir los archivos
set "files=cmd\main.go config\config.go internal\database\db.go internal\database\migration.go internal\database\models.go internal\handlers\chatbot.go internal\handlers\history.go internal\services\chatbot.go pkg\logger.go scripts\start.bat go.mod go.sum README.md"

:: Crear carpetas
for %%F in (%folders%) do (
    if not exist %%F mkdir %%F
)

:: Crear archivos vacíos
for %%F in (%files%) do (
    if not exist %%F type nul > %%F
)

echo Estructura de proyecto creada correctamente.
exit
