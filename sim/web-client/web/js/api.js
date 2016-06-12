/**
 * APIアクセスクラス
 */
class API {
    /**
     * constructor
     */
    constructor() {
        this.baseUrl = "http://localhost:1323/";
    }

    /**
     * メッセージ一覧取得
     * GET: messages
     *
     * @param success
     * @param error
     * @returns {*}
     */
    getMessages(success, error) {
        var getMessageUri = this.baseUrl + "messages";
        return $.ajax({
                type: "get",
                url: getMessageUri
            })
            .done(function(data) { success(data) })
            .fail(function() { error() });
    }

    /**
     * メッセージ投稿
     * POST: messages
     *
     * @param body
     * @param success
     * @param error
     * @returns {*}
     */
    postMessage(body, success, error) {
        var postMessageUri = this.baseUrl + "messages";
        return $.ajax({
            type: "post",
            url: postMessageUri,
            contentType: "application/json",
            data: JSON.stringify({body:body})
        })
        .done(function(data) { success(); console.log(data) })
        .fail(function() { error() });
    }

    /**
     * メッセージ更新
     * POST: messages/[id]
     * TODO: サーバ側で実装されていません
     *
     * @param id
     * @param body
     * @param success
     * @param error
     * @returns {*}
     */
    updateMessage(id, body, success, error) {
        var postMessageUri = this.baseUrl + "messages/" + id;
        return $.ajax({
                type: "put",
                url: postMessageUri,
                contentType: "application/json",
                data: JSON.stringify({body:body})
            })
            .done(function(data) { success(); console.log(data) })
            .fail(function() { error() });
    }

    /**
     * メッセージ更新
     * POST: messages/[id]
     * TODO: サーバ側で実装されていません
     *
     * @param id
     * @param success
     * @param error
     * @returns {*}
     */
    deleteMessage(id, success, error) {
        var postMessageUri = this.baseUrl + "messages/" + id;
        return $.ajax({
                type: "delete",
                url: postMessageUri,
                contentType: "application/json"
            })
            .done(function(data) { success(); console.log(data) })
            .fail(function() { error() });
    }
}
