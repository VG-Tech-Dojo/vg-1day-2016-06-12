//
//  MassageTableViewCell.swift
//  OneDayInternshipApp
//
//  Created by 清 貴幸 on 2016/05/26.
//  Copyright © 2016年 VOYAGE GROUP, Inc. All rights reserved.
//

import UIKit

protocol MessageTableViewCellDelegate {
    func didLongPressCell(cell: MessageTableViewCell)
}

class MessageTableViewCell: UITableViewCell {
    @IBOutlet weak private var messageLabel: UILabel!
    @IBOutlet weak var datetimeLabel: UILabel!
    @IBOutlet weak var usernameLabel: UILabel!
    
    var delegate: MessageTableViewCellDelegate? = nil
    
    override func awakeFromNib() {
        self.addGestureRecognizer(UILongPressGestureRecognizer(target: self, action: #selector(handleLongPress)))
    }
    
    override func prepareForReuse() {
        self.messageLabel.text = nil
        self.datetimeLabel.text = nil
        self.usernameLabel.text = nil
    }
    
    func setupComponentsWithMessage(message: Message) {
        self.messageLabel.text = message.body
        self.datetimeLabel.text = message.date
        self.usernameLabel.text = message.username
    }
    
    func handleLongPress(recognizer: UILongPressGestureRecognizer) {
        self.delegate?.didLongPressCell(self)
    }
}
