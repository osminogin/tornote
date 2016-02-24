$(document).ready(function() {
    $("#main").removeClass("hidden");
    $("#warn").remove();

    // Submit new note
    $("#note").submit(function(event) {
        var form = $(this);
        var text = form.find("textarea").val();
        var secret = CryptoJS.lib.WordArray.random(128/8);
        var encrypted = CryptoJS.AES.encrypt(text, secret);

        $.ajax({
            url: form.attr("action"),
            method: "POST",
            data: {body: encodeURIComponent(encrypted)},
            success: function(data) {
                $("#note").addClass("hidden");
                $("#done").removeClass("hidden");
                // XXX:
                window.alert(data);
            }
        });
        event.preventDefault();
    });

    // Display link

    // Show secret note
    //var decrypted = CryptoJS.AES.decrypt(encrypted, "Secret Passphrase");
});
