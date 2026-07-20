#!/usr/bin/env bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}==========================================${NC}"
echo -e "${BLUE}       Building Kalkulator Rowków         ${NC}"
echo -e "${BLUE}==========================================${NC}"

# Check if wails is installed
if ! command -v wails &> /dev/null; then
    echo -e "${RED}Error: 'wails' CLI is not installed or not in PATH.${NC}"
    echo "Please install Wails by running: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    exit 1
fi

CLEAN_FLAG=""
if [ "$1" == "--clean" ] || [ "$1" == "-clean" ]; then
    CLEAN_FLAG="-clean"
    echo -e "${BLUE}Clean flag detected. Bin directory will be cleaned before building.${NC}"
    rm -rf build/bin/*
fi

# Sync icon from app-icon.png in main folder
if [ -f "app-icon.png" ]; then
    echo -e "${BLUE}Updating app icon assets from root app-icon.png...${NC}"
    cp app-icon.png build/appicon.png
    cp app-icon.png frontend/src/assets/images/logo-universal.png
fi

# 1. Build for macOS (Universal binary: Apple Silicon + Intel)
echo -e "\n${GREEN}[1/2] Building for macOS (darwin/universal)...${NC}"
wails build -platform darwin/universal -o KalkulatorRowkow $CLEAN_FLAG

# 2. Build for Windows (x64)
echo -e "\n${GREEN}[2/2] Building for Windows (windows/amd64)...${NC}"
wails build -platform windows/amd64 -o KalkulatorRowkow.exe

echo -e "\n${BLUE}==========================================${NC}"
echo -e "${GREEN}✓ Build process completed successfully!${NC}"
echo -e "${BLUE}Outputs created in: build/bin/${NC}"
echo -e "${BLUE}==========================================${NC}"
ls -lh build/bin
