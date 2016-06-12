//
//  Endpoint.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/31.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import Foundation

let baseURLString = "http://localhost:1323/"

enum Endpoint: String {
    case Messages = "messages"
    
    var URLString: String {
        switch self {
        case .Messages:
            return baseURLString + self.rawValue
        }
    }
    
    var URL: NSURL {
        return NSURL(string: self.URLString)!
    }
    
    func URLWithPathString(pathString: String) -> NSURL {
        return NSURL(string: self.URLString + "/" +
            pathString)!
    }
}
