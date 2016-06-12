//
//  Message.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/26.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import Foundation

class Message {
    let identifier: Int?
    let body: String!
    // 1-1 時刻データを扱うためのプロパティ追加
    let date: String
    
    // 1-2 ユーザ名データを扱うためのプロパティ追加
    
    init(identifier: Int?, body: String, date: String) {
        self.identifier = identifier
        self.body = body
        // 1-1 時刻データを用いて初期化
        self.date = date
        
        // 1-2 ユーザ名データを用いて初期化
    }
}
