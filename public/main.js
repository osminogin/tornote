$(document).ready(function() {
    $("#main").removeClass("hidden");
    $("#warn").remove();

    // Submit new secret note
    $("#note").submit(function(event) {
        event.preventDefault();
        var form = $(this);
        var text = form.find("textarea").val();
        var secret = sjcl.codec.base64url.fromBits(sjcl.random.randomWords(3));
        var encrypted = sjcl.encrypt(secret, text);

        $.ajax({
            url: form.attr("action"),
            method: "POST",
            data: {body: encrypted.toString()},
            success: function(id) {
                var link = window.location.href.toString() + id + "#" + secret;
                $("#secret_link").text(link);
                $("#note").addClass("hidden");
                $("#done").removeClass("hidden");
            }
        });
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
