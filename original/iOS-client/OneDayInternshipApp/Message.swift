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
    let date: String
    let username: String
    
    init(identifier: Int?, body: String, date: String, username: String) {
        self.identifier = identifier
        self.body = body
        self.date = date
        self.username = username
    }
}
