#!/bin/bash

src_dir="/home/zks/Documents/video"

# 遍历目录下的所有.mp4文件
for mp4_file in "$src_dir"/*.mp4; do
    # 从.mp4文件名中提取文件名部分（不包含扩展名）
    mp4_base=$(basename "$mp4_file" .mp4)
    
    # 构造对应的.png文件路径
    png_file="$src_dir/$mp4_base.png"
    
    # 如果对应的.png文件存在，则重命名
    if [ -e "$png_file" ]; then
        mv "$png_file" "$src_dir/$mp4_base.mp4.png"
    fi
done
