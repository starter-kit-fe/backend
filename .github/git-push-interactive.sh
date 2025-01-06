#!/bin/bash

# 暂存所有更改
git add .

# 提示用户输入提交信息
read -p "Enter commit message: " COMMIT_MSG

# 检查提交信息是否为空
if [ -z "$COMMIT_MSG" ]; then
    echo "Commit message cannot be empty. Aborting."
    exit 1
fi

# 执行 Git 提交
git commit -m "$COMMIT_MSG"

# 推送到远程仓库
git push

