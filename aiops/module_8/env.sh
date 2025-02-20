# Fix the initial PATH setup
export PATH="/usr/local/bin:$PATH"  # 将 /usr/local/bin 添加到 PATH 的开头
export PATH="/sbin:/usr/bin:/bin:/usr/sbin:$PATH"  # 添加其他基本路径到 PATH

# Set Go environment variables
export GOPATH="/Users/chenchen/Documents/Go"
export GOROOT="/usr/local/go"

# Append Go and Docker paths
export PATH="$PATH:$GOPATH/bin:$GOROOT/bin"
export PATH="$PATH:/usr/local/bin/docker"
export PATH="$PATH:/usr/local/bin/gitlab-runner"

# Add Allure, Python, Google Cloud SDK, and Framework paths
export PATH="$PATH:/Applications/Visual Studio Code.app/Contents/Resources/app/bin"
export PATH="$PATH:/Users/chenchen/Applications/allure-2.19.0/bin"
export PATH="$PATH:/Library/Frameworks/Python.framework/Versions/3.12/bin"
export PATH="$PATH:/Users/chenchen/Downloads/google-cloud-sdk/bin"
export PATH="$PATH:/Users/chenchen/Desktop/automatedtest/framework"

# Aliases
alias cobra-cli="~/go/bin/cobra-cli"
alias python="/Library/Frameworks/Python.framework/Versions/3.12/bin/python3"
alias pip="/Library/Frameworks/Python.framework/Versions/3.12/bin/pip3.12"
alias chen='ssh chen@3.13.3.30'

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/Users/chenchen/Downloads/google-cloud-sdk/path.zsh.inc' ]; then . '/Users/chenchen/Downloads/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/Users/chenchen/Downloads/google-cloud-sdk/completion.zsh.inc' ]; then . '/Users/chenchen/Downloads/google-cloud-sdk/completion.zsh.inc'; fi

# NVM configuration
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# GVM configuration
[[ -s "/Users/chenchen/.gvm/scripts/gvm" ]] && source "/Users/chenchen/.gvm/scripts/gvm"
[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
