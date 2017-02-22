package mail

var accountResetPasswordEmailBody = `
<!DOCTYPE html>
<html>
    <body style="background-color:white">
        <div style="color:#404040;">
         	<h2>Hi <b style="color:red">{accountName}</b>!</h2>
            <h3>If you requested a new password for your AMP account, please use the following command line:</h4>
            <h4>amp account password --set {token}</h4>
        <div style="height:20px"></div>
        <div style="color:#404040;">
            <h4>If you didn't make this request, you can safely ignore this email</h4>
            <h4>This link is good for one hour only</h4>
        </div>
    </body>
</html>

`

/* possible images
   <img atl="" src="data:image/jpeg;base64,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"/>
   <img alt="" src="data:image/jpeg;base64,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"/>
*/
