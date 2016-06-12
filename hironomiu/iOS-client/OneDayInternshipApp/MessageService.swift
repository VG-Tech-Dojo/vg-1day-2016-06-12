//
//  MessageService.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/31.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import Foundation

class MessageService {
    static let sharedService = MessageService()
    
    private init() {
    }
    
    var messages: [Message] = []
    
    func fetch(completionHandler: (NSError?) -> Void) {
        
        let handler: (NSData?, NSURLResponse?, NSError?) -> Void = {
            [weak self] (data, response, error) in
            if error != nil {
                print("error: \(error?.domain)")
                completionHandler(error)
            }
            
           assert(data != nil, "No data response from server")
            
            do {
                if let JSONObject: AnyObject! = try NSJSONSerialization.JSONObjectWithData(data!, options: NSJSONReadingOptions.MutableContainers), messages: [Message] = self?.parseJSONObjectToMessageArray(JSONObject) {
                    self?.messages = messages
                }
            } catch let error as NSError {
                print("error: \(error.domain)")
            }
            
            completionHandler(error)
        }
        
        APIRequest.HTTPRequestWithURL(Endpoint.Messages.URL, HTTPMethod: "GET", JSONDictionary: nil, completionHandler: handler)
    }
    
    // 1-2 ユーザ名データを投稿できるよう修正
    static func post(messageBody: String, completionHandler: NSError? -> Void) {
        APIRequest.HTTPRequestWithURL(Endpoint.Messages.URL, HTTPMethod: "POST", JSONDictionary: ["body": messageBody]) { (data, response, error) in
                completionHandler(error)
        }
    }

    
    static func update(messageBody: String, messageID: Int, completionHandler: NSError? -> Void) {
        APIRequest.HTTPRequestWithURL(Endpoint.Messages.URLWithPathString("\(messageID)"), HTTPMethod: "PUT", JSONDictionary: ["body": messageBody]) { (data, response, error) in
            completionHandler(error)
        }
    }
    
    static func delete(messageID: Int, completionHandler: NSError? -> Void) {
        APIRequest.HTTPRequestWithURL(Endpoint.Messages.URLWithPathString("\(messageID)"), HTTPMethod: "DELETE", JSONDictionary: ["id": messageID]) { (data, response, error) in
            completionHandler(error)
        }
    }
    
    private func parseJSONObjectToMessageArray(JSONObject: AnyObject!) -> [Message] {
        var messages: [Message] = []
        
        guard let JSONArray = JSONObject as? [[String: AnyObject]] else {
            return messages
        }
        
        for json in JSONArray {
            // 1-1, 1-2 新規で追加した値のパース処理を追加
            let message = Message(identifier: json["id"] as! Int?, body: json["body"] as! String)
            messages.append(message)
        }
        
        return messages
    }
}