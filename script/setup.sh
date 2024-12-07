#!/bin/zsh

set -e

echo "Atualizando pacotes do sistema..."
sudo apt update && sudo apt upgrade -y
sudo apt-get install build-essential

# 1. Verificar se o Go já está instalado
if command -v go &> /dev/null; then
    echo "Go já está instalado. Removendo a instalação anterior..."
    sudo chmod -R 777 $HOME/go
    sudo chown -R $USER:$USER $HOME/go
    rm -rf $HOME/go
    sudo rm -rf /usr/local/go
else
    echo "Go não está instalado."
fi

# 2. Baixar o Go versão 1.23.4
echo "Baixando o Go versão 1.23.4..."
curl -LO https://go.dev/dl/go1.23.4.linux-amd64.tar.gz

echo "Extraindo o Go..."
# Extrair o Go na pasta do usuário
tar -C $HOME -xzf go1.23.4.linux-amd64.tar.gz

# Remover o arquivo de instalação
rm go1.23.4.linux-amd64.tar.gz

# 3. Configurar o Go no PATH
echo "Configurando o PATH para Go..."
echo "export GOPATH=$HOME/go" >> ~/.zshrc
echo "export GOROOT=$GOPATH" >> ~/.zshrc
echo "export PATH=\$PATH:\$GOROOT/bin" >> ~/.zshrc
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.zshrc
source ~/.zshrc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

echo "Go versão 1.23.4 instalado com sucesso: $(go version)"

# 4. Instalar o Homebrew sem interação
echo "Instalando o Homebrew..."
NONINTERACTIVE="1" /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Adicionar Homebrew ao PATH
echo "Adicionando Homebrew ao PATH..."
echo "eval \"\$($HOME/.linuxbrew/bin/brew shellenv)\"" >> ~/.zshrc
source ~/.zshrc

echo "Homebrew instalado com sucesso: $(brew --version)"

# 5. Remover o `buf` caso esteja instalado
if command -v buf &> /dev/null; then
    echo "buf já está instalado. Removendo a versão existente..."
    brew uninstall buf
fi

# 6. Instalar a última versão do `protoc`
echo "Baixando o Protocol Buffers (protoc)..."
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v21.10/protoc-21.10-linux-x86_64.zip

echo "Extraindo o protoc..."
unzip protoc-21.10-linux-x86_64.zip -d $HOME/protoc

# Adicionar ao PATH
echo "Adicionando o `protoc` ao PATH..."
echo "export PATH=\$PATH:$HOME/protoc/bin" >> ~/.zshrc
source ~/.zshrc

# Remover o arquivo de instalação
rm protoc-21.10-linux-x86_64.zip

echo "Protocol Buffers (protoc) instalado com sucesso: $(protoc --version)"

echo "Script concluído com sucesso!"
