//
//  PostViewController.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/27.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import UIKit

protocol PostViewControllerDelagate : NSObjectProtocol {
    func postViewController(viewController : PostViewController, didTouchUpCloseButton: AnyObject)
}

class PostViewController: UIViewController, UITextViewDelegate {
    @IBOutlet weak private var messageTextView: UITextView!
    weak var delegate: PostViewControllerDelagate?
    @IBOutlet weak var sendButton: UIButton!
    var isUpdate: Bool = false
    var message: Message?
    @IBOutlet weak var usernameTextField: UITextField!

    override func viewDidLoad() {
        super.viewDidLoad()
        self.messageTextView.becomeFirstResponder()
        
        if self.message?.body.characters.count > 0 {
            self.messageTextView.text = self.message?.body
        }
    }
    
    func postAndCloseWithMessageBody(messageBody: String, sender: AnyObject) {
        
        // 1-2 ユーザ名データを投稿できるよう修正
        MessageService.post(messageBody) {
            [weak self] (error) in
            
            self?.delegate?.postViewController(self!, didTouchUpCloseButton: sender)
            if error != nil {
                // TODO: エラー処理
                print("Error: \(error?.domain)")
                return
            }
        }
    }
    
    func updateAndCloseWithMessageBody(messageBody: String, messageID: Int, sender: AnyObject) {
        MessageService.update(messageBody, messageID: messageID) {
            [weak self] (error) in
            
            self?.delegate?.postViewController(self!, didTouchUpCloseButton: sender)
            if error != nil {
                // TODO: エラー処理
                print("Error: \(error?.domain)")
                return
            }
        }
    }
    
    // MARK: - IBAction
    
    @IBAction func didTouchUpCloseButton(sender: AnyObject) {
        self.messageTextView.resignFirstResponder()
        self.delegate?.postViewController(self, didTouchUpCloseButton: sender)
    }
    
    @IBAction func didTouchUpSendButton(sender: AnyObject) {
        self.messageTextView.resignFirstResponder()
        let messageBody = self.messageTextView.text ?? ""

        if self.isUpdate {
            self.updateAndCloseWithMessageBody(messageBody, messageID: (self.message?.identifier)!, sender: sender)
            return
        }
        
        self.postAndCloseWithMessageBody(messageBody, sender: sender)
    }
    
    // MARK: - UITextViewDelegate
    
    func textViewDidChange(textView: UITextView) {
        if textView.text.characters.count > 0 {
            self.sendButton.enabled = true
            return
        }
        
        self.sendButton.enabled = false
    }
    
    func textView(textView: UITextView, shouldChangeTextInRange range: NSRange, replacementText text: String) -> Bool {
        if text == "\n" {
            textView.resignFirstResponder()
            self.didTouchUpSendButton(self)
            return false
        }
        
        return true
    }
}
