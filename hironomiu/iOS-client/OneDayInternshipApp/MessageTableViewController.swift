//
//  MessageTableViewController.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/26.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import UIKit

class MessageTableViewController: UITableViewController, PostViewControllerDelagate, MessageTableViewCellDelegate {
    var messages: [Message] = []
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.setupRefreshControl()
    }
    
    override func viewDidAppear(animated: Bool) {
        super.viewDidAppear(animated)
        self.fetchMessagesAndRefresh()
    }
    
    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
        let destVC: PostViewController = segue.destinationViewController as! PostViewController
        destVC.delegate = self

        if segue.identifier == "editSegue" {
            destVC.message = self.messages[self.tableView.indexPathForSelectedRow!.row]
        }
    }
    
    func setupRefreshControl() {
        let refreshControl = UIRefreshControl()
        refreshControl.attributedTitle = NSAttributedString(string: "引っ張って更新")
        refreshControl.addTarget(self, action: #selector(MessageTableViewController.fetchMessagesAndRefresh), forControlEvents: UIControlEvents.ValueChanged)
        self.refreshControl = refreshControl
    }
    
    func fetchMessagesAndRefresh() {
        MessageService.sharedService.fetch {
            [weak self] (error) in
            
            if error != nil {
                print("error: \(error?.domain)")
                return
            }
            
            self?.messages = MessageService.sharedService.messages
            
            dispatch_async(dispatch_get_main_queue(), {
                self?.refreshControl?.endRefreshing()
                self?.tableView.reloadData()
            })
        }

        self.refreshControl?.beginRefreshing()
    }

    
    func deleteMessage(index: Int) {
        MessageService.delete(index) { (error) in
            if error != nil {
                print("error: \(error?.domain)")
                return
            }
        }
    }
    
    // MARK: - UITableViewDataSource
    
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.messages.count
    }
    
    // MARK: - UITableViewDelegate

    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("MessageTableViewCell", forIndexPath: indexPath) as! MessageTableViewCell
        
        let message = self.messages[indexPath.row]
        cell.setupComponentsWithMessage(message)
        cell.delegate = self
        
        return cell
    }
    
    override func tableView(tableView: UITableView, canEditRowAtIndexPath indexPath: NSIndexPath) -> Bool {
        return true
    }
    
    override func tableView(tableView: UITableView, commitEditingStyle editingStyle: UITableViewCellEditingStyle, forRowAtIndexPath indexPath: NSIndexPath) {
        let message = self.messages[indexPath.row]
        self.messages.removeAtIndex(indexPath.row)
        self.tableView.deleteRowsAtIndexPaths([indexPath], withRowAnimation: .Fade)
        
        self.deleteMessage(message.identifier!)
    }
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        tableView.deselectRowAtIndexPath(indexPath, animated: true)
    }
    
    // MARK: - MessageTableViewCellDelegate
    
    func didLongPressCell(cell: MessageTableViewCell) {
        let index = self.tableView.indexPathForCell(cell)!.row
        self.tableView.selectRowAtIndexPath(NSIndexPath(indexes: [0, index], length: 2), animated: false, scrollPosition: .None)
        self.performSegueWithIdentifier("editSegue", sender: self)
    }
    
    // MARK: - IBAction
    
    func postViewController(viewController: PostViewController, didTouchUpCloseButton: AnyObject) {
        self.presentedViewController?.dismissViewControllerAnimated(true, completion: nil)
    }
}
