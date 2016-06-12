//
//  APIRequest.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/26.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import Foundation

class APIRequest {
    static func HTTPRequestWithURL(URL: NSURL, HTTPMethod: String, JSONDictionary: Dictionary<String, AnyObject>?, completionHandler: (NSData?, NSURLResponse?, NSError?) -> Void) {
        var body: NSData? = nil;
        
        do {
            if let dic = JSONDictionary {
                body = try NSJSONSerialization.dataWithJSONObject(dic, options: NSJSONWritingOptions.PrettyPrinted)
            }
        } catch let error as NSError {
            print("Error: \(error.domain)")
            completionHandler(nil, nil, error)
        }
        
        let request = self.URLRequest(URL, HTTPMethod: HTTPMethod, HTTPBody: body)
        let session = NSURLSession.sharedSession()
        let task = session.dataTaskWithRequest(request, completionHandler: completionHandler)
        task.resume()
    }
    
    private static func URLRequest(URL: NSURL, HTTPMethod: String, HTTPBody: NSData?) -> NSURLRequest {
        let request = NSMutableURLRequest(URL: URL)
        request.HTTPMethod = HTTPMethod
        request.HTTPBody = HTTPBody
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        return request.copy() as! NSURLRequest
    }}