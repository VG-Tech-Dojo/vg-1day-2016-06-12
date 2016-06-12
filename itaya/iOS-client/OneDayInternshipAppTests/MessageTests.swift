//
//  MessageTests.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/26.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import XCTest

class MessageTests: XCTestCase {
    
    func testInit() {
        let fetchdMessage = Message(identifier: 0, body: "取得されたメッセージ")
        XCTAssertEqual(fetchdMessage.body, "取得されたメッセージ")
        XCTAssertEqual(fetchdMessage.identifier, 0)
        
        let messageForPost = Message(identifier: nil, body: "投稿用")
        XCTAssertEqual(messageForPost.identifier, nil)
        XCTAssertEqual(messageForPost.body, "投稿用")
    }
    
}
