/**
 * イベントが発生した時に呼ばれる関数
 */

/**
 * メッセージリストをHTMLに挿入
 * メッセージの更新や読み込みが完了した場合に呼ばれる
 *
 * @param data
 */
function appendMessages(data) {
    $("#message-container").empty();
    for ( var i = 0; i < data.length; i++ ) {
        var object = data[i];
        appendMessage(object);
    }
}

/**
 * メッセージ１件をHTMLに挿入
 * Todo:メッセージの表示形式を変更する場合はここを修正する
 *
 * @param message
 */
function appendMessage(message) {
    // Bodyをエスケープ
    var escapeBody = $("<div/>").text(message.body).html();
    var messageHTML =
        '<div class="media">' +
            '<div class="media-body">' +
                //'<span class="media-message-name">名無しさん</span>  ' +
                //'<span class="media-message-date">' + escapeDate + '</span>' + '<br>' +
                '<span class="media-message-body">' + escapeBody + '</span>' +
            '</div>' +
            '<div class="media-right">' +
                '<button type="button" class="pull-right btn btn-default btn-xs" data-toggle="modal" data-target="#edit-modal" data-body="' + escapeBody + '" data-id="' + message.id +'">' +
                '<span class="glyphicon glyphicon-option-horizontal" aria-hidden="true"></span>' +
            '</div>' +
        '</div>' +
        '<hr>';

    $("#message-container").append(messageHTML);
}

/**
 * メッセージリストの再読み込み
 */
function reloadMessages() {
    var success = function(data) {
        appendMessages(data);
    };
    var error = function() { console.log("error") };
    (new API()).getMessages(success, error);
}

/**
 * メッセージの投稿
 *
 * @param body
 */
function sendMessage(body) {
    var success = function() {
        $(".message-body").val("");
        reloadMessages();
    };
    var error = function() { console.log("error") };
    (new API()).postMessage(body, success, error);
}

/**
 * メッセージの更新
 *
 * @param id
 * @param body
 */
function updateMessage(id, body) {
    var success = function() {
        reloadMessages();
    };
    var error = function() { console.log("error") };
    (new API()).updateMessage(id, body, success, error);
}

/**
 * メッセージの削除
 *
 * @param id
 */
function deleteMessage(id) {
    var success = function() {
        reloadMessages();
    };
    var error = function() { console.log("error") };
    (new API()).deleteMessage(id, success, error);
}


