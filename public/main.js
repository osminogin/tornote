// Copyright 2016-2020 Vladimir Osintsev <osintsev@gmail.com>
//
// This file is part of Tornote.
//
// Tornote is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Tornote is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

$(document).ready(function() {
    $("#main").removeClass("hidden");
    $("#warn").remove();

    // Submit new secret note
    $("#note").submit(function(event) {
        let form = $(this);
        let text = form.find("textarea").val();
        let secret = sjcl.codec.base64url.fromBits(sjcl.random.randomWords(5));
        let encrypted = sjcl.encrypt(secret, text);
        let csrfToken = document.getElementsByName("csrf_token")[0].value;

        $.ajax({
            url: form.attr("action"),
            method: "POST",
            xhrFields: {
                withCredentials: true
            },
            data: {
                body: encrypted.toString(),
            },
            headers: {"X-CSRF-Token": csrfToken},
            success: function(id) {
                let link = window.location.href.toString() + id + "#" + secret;
                $("#secret_link").text(link);
                $("a", "#done").first().attr("href", link);
                $("#note").addClass("hidden");
                $("#done").removeClass("hidden");
                SelectText("secret_link");
            },
            error: function (err) {
                window.alert(err.responseText);
            }
        });
        event.preventDefault();
    });

    // Copy to buffer
    $("a.btn-primary").click(function(event) {
        // TODO: Show feedback to user
        document.execCommand('copy');
        event.preventDefault();
    });

    // Show decrypted secret note
    if($("#secret_note").length > 0){
        let secret = window.location.hash.substring(1);
        let cipherText = $("#secret_note").text();
        let decrypted = sjcl.decrypt(secret, cipherText);
        $("#secret_note").html(decrypted);
        $("#secret_note").removeClass("hidden");
    }
});

// Soluiton from https://stackoverflow.com/questions/985272/
function SelectText(element) {
    let doc = document,
        text = doc.getElementById(element),
        range,
        selection;
    if (doc.body.createTextRange) {
        range = document.body.createTextRange();
        range.moveToElementText(text);
        range.select();
    } else if (window.getSelection) {
        selection = window.getSelection();
        range = document.createRange();
        range.selectNodeContents(text);
        selection.removeAllRanges();
        selection.addRange(range);
    }
}