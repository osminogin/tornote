// Copyright 2020 Vladimir Osintsev <osintsev@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

$(document).ready(function() {
    $("#main").removeClass("hidden");
    $("#warn").remove();

    // Submit new secret note
    $("#note").submit(function(event) {
        var form = $(this);
        var text = form.find("textarea").val();
        var secret = sjcl.codec.base64url.fromBits(sjcl.random.randomWords(5));
        var encrypted = sjcl.encrypt(secret, text);

        $.ajax({
            url: form.attr("action"),
            method: "POST",
            data: {body: encrypted.toString()},
            success: function(id) {
                var link = window.location.href.toString() + id + "#" + secret;
                $("#secret_link").text(link);
                $("a", "#done").first().attr("href", link);
                $("#note").addClass("hidden");
                $("#done").removeClass("hidden");
                SelectText("secret_link");
            }
        });
        event.preventDefault();
    });

    // Show decrypted secret note
    if($("#secret_note").length > 0){
        var secret = window.location.hash.substring(1);
        var ciphertext = $("#secret_note").text();
        var decrypted = sjcl.decrypt(secret, ciphertext);
        $("#secret_note").html(decrypted);
        $("#secret_note").removeClass("hidden");
    };

});

// Soluiton from https://stackoverflow.com/questions/985272/
function SelectText(element) {
    var doc = document,
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