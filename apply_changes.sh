#!/bin/bash

# ^\d+\.\d+$ にマッチするブランチ名を取得しループ処理
git branch --list | grep '^..[0-9]\+\.[0-9]\+$' | while read -r branch; do
    echo ""
    echo "-------- $branch --------"

    # ブランチ名の前にある空白をトリミング
    branch=$(echo "$branch" | xargs)

    # ブランチをチェックアウト
    git checkout "$branch"

    # masterブランチの修正を取り込む
    git merge master

    # 変更をpushする
    git push origin $branch

    # masterブランチをチェックアウト
    git checkout master
done
