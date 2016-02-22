//var decrypted = CryptoJS.AES.decrypt(encrypted, "Secret Passphrase");
$(document).ready(function() {
    $("#main").removeClass("hidden");
    $("#warn").remove();
    $("#note").submit(function(event) {
        var data = "aaa";
        var secret = "1234567890";
        var encrypted = CryptoJS.AES.encrypt(data, secret);
        $.post( "/note/aaa", {body: encrypted}, function(data) {
                window.alert(data);
        });
        event.preventDefault();
        event.unbind();
    });
});
