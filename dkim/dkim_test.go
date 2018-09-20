package dkim

import (
	"log"
	"strings"
	"testing"

	dkim "github.com/emersion/go-dkim"
)

func TestDKIM(t *testing.T) {
	r := strings.NewReader(mailString)

	verifications, err := dkim.Verify(r)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range verifications {
		if v.Err == nil {
			log.Println("Valid signature for:", v.Domain)
		} else {
			log.Println("Invalid signature for:", v.Domain, v.Err)
		}
	}

}

const mailString = `Return-Path: <107-FMS-070.0.8007.0.0.4076.9.8319249@em.oreilly.com>
Delivered-To: mlecarme@bearstech.com
Received-SPF: Pass (sender SPF authorized) identity=mailfrom; client-ip=192.28.152.137; helo=em.oreilly.com; envelope-from=107-fms-070.0.8007.0.0.4076.9.8319249@em.oreilly.com; receiver=mlecarme@bearstech.com 
Received: from em.oreilly.com (em.oreilly.com [192.28.152.137])
	(using TLSv1.2 with cipher ECDHE-RSA-AES256-GCM-SHA384 (256/256 bits))
	(No client certificate requested)
	by bernardo.bearstech.com (Postfix) with ESMTPS id 0CA5214201BE
	for <mlecarme@bearstech.com>; Wed, 19 Sep 2018 17:32:16 +0200 (CEST)
X-MSFBL: bWxlY2FybWVAYmVhcnN0ZWNoLmNvbUBkdnAtMTkyLTI4LTE1Mi0xMzdAYmctYWJk
	LTYyMkAxMDctRk1TLTA3MDoyNTk3OjYyMTk6MTk0NDU6MDo0MDc2Ojk6ODAwNzo4
	MzE5MjQ5
Received: from [10.1.87.249] ([10.1.87.249:37968] helo=abmas01.marketo.org)
	by abmta22.marketo.org (envelope-from <reply@oreilly.com>)
	(ecelerity 3.6.8.47404 r(Core:3.6.8.0)) with ESMTP
	id 46/8C-18300-DFB62AB5; Wed, 19 Sep 2018 10:32:13 -0500
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/relaxed; t=1537371133;
	s=m1; d=oreilly.com; i=@oreilly.com;
	h=Date:From:To:Subject:MIME-Version:Content-Type;
	bh=Ge8M/7pe1mU2DX9rpRpjv9LXCv1/nK3dG53MoZ4l2SQ=;
	b=JIJwaurGijua2LaBe+Xor4amj+PjM2D4s6yct6XexYidNZHsacny9mI9vBs+GHUq
	krDrRRSdfvit/cWPuK5YuBO18rmb7INE7pyOsT8xprbjCBDUPh/lvMDanOwIw6qC40h
	yX5BqTyjIakYjoR2G2DWX52k85EmcxiWrfJ6PU1Q=
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/relaxed; t=1537371133;
	s=m1; d=mktdns.com; i=@mktdns.com;
	h=Date:From:To:Subject:MIME-Version:Content-Type;
	bh=Ge8M/7pe1mU2DX9rpRpjv9LXCv1/nK3dG53MoZ4l2SQ=;
	b=hylmOWEdXKcmIjEVp4rzKhYsIaCdi1wa1ydabVP/vkdv3CUXqnkdLwsnsvS7scaM
	aQOitV0vtfE6hlbWnrGvokNcOKqultt8Fx2ae92m9dbC9ghfWEPiDVGwIGOW3Gl0fE1
	kd9HHi/P3yJxBkVdd+lcBm5xS0q2CoHhP2APQar0=
Date: Wed, 19 Sep 2018 10:32:13 -0500 (CDT)
From: O'Reilly Velocity Conference <reply@oreilly.com>
Reply-To: reply@oreilly.com
To: mlecarme@bearstech.com
Message-ID: <306549440.-1790310584.1537371133215.JavaMail.root@abmas01.marketo.org>
Subject: Get Kubernetes certified in just 2 days
MIME-Version: 1.0
Content-Type: multipart/alternative; 
	boundary="----=_Part_-1790310585_179862840.1537371133215"
X-Binding: bg-abd-622
X-MarketoID: 107-FMS-070:2597:6219:19445:0:4076:9:8007:8319249
X-MktArchive: false
List-Unsubscribe: <mailto:M5SUQTRTJRTGCTSWMRQVS6DXKVTC2UJVNB3T2PI.8007.4076.9@unsub-ab.mktomail.com>
X-Mailfrom: 107-FMS-070.0.8007.0.0.4076.9.8319249@em.oreilly.com
X-MSYS-API: {"options":{"open_tracking":false,"click_tracking":false}}
X-MktMailDKIM: true
X-Virus-Scanned: clamav-milter 0.100.1 at bernardo
X-Virus-Status: Clean

------=_Part_-1790310585_179862840.1537371133215
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 7bit

View this information in your browser: https://get.oreilly.com/index.php/email/emailWebview?mkt_tok=eyJpIjoiTURNNE5tUmhPVEZqT0dZMyIsInQiOiJlaUJTSCtPN1o0dUh3bzd4YTRpN3I5OTNqckw5bTNTQWdzcVFQcEJJQmd3S1RRQXprS0hhNGJmekZpa2JXWW9iNlZocXlaSVgzVys4ZGNnakdmQVpROHpUS1BUU1BETWZ6UzFaTG5IZkFTN2JxK3hqandzZ3lyYm9IbXNGY1ZVQSJ9

Hi Mathieu,

There's an important question that many hiring managers are asking software engineer candidates more often:

"Can you develop and maintain applications using Kubernetes?"

What's more, according to the 2017 Linux Foundation and Dice Open Source Jobs Report, (https://www.linuxfoundation.org/publications/2017/08/open-source-jobs-report-2017/?utm_medium=email&utm_source=topic+optin&utm_campaign=vleu18&utm_content=20180919+08a+ckad+training+full) half of hiring managers prefer to hire candidates with certifications, 65% seek expertise in Linux, and 70% seek expertise in cloud technologies.

That's where the Certified Kubernetes Application Developer (CKAD) prep course and exam comes in. This 2-day program at the O'Reilly Velocity Conference is the perfect way to increase your marketability and value in the job market, putting you ahead of the pack as this emerging platform grows in popularity.

The CKAD prep course and exam will give you:

- An overview of the key concepts and skills you need to know to pass the official Certified Kubernetes Application Developer (CKAD) exam

- A review of the core API primitives and basic security concepts such as service accounts and Pod security policies

- An understanding of all Pod concepts and API schema, how to configure applications, and Pod monitoring and logging practices

Most importantly, you can leave Velocity a Certified Kubernetes Application Developer. Just pick the date and location that works best for you: September 30-October 1 in New York City, or October 30-31 in London.

Register for Velocity in New York: https://conferences.oreilly.com/velocity/vl-ny/public/schedule/detail/70472?utm_medium=email&utm_source=topic+optin&utm_campaign=vleu18&utm_content=20180919+08a+ckad+training+full

Register for Velocity in London: https://conferences.oreilly.com/velocity/vl-eu/public/schedule/detail/71386?utm_medium=email&utm_source=topic+optin&utm_campaign=vleu18&utm_content=20180919+08a+ckad+training+full

See you at Velocity,

The O'Reilly Team

------------------------------------------------------

Not planning to attend Velocity in 2018? Pause these emails until next year:
http://www.oreilly.com/emails/velocity-subscription.html?utm_medium=email&utm_source=topic+optin&utm_campaign=vleu18&utm_content=20180919+08a+ckad+training+full

This message was sent to mlecarme@bearstech.com. You are receiving this because you're a customer of O'Reilly Media, or you've signed up to receive email from us. We hope you found this message to be useful. However, if you'd rather not receive future emails of this type from O'Reilly, please manage your preferences or unsubscribe here:
https://get.oreilly.com/email-preferences.html?utm_medium=email&utm_source=topic+optin&utm_campaign=vleu18&utm_content=20180919+08a+ckad+training+full

Please read our Privacy Policy: https://www.oreilly.com/privacy.html

O'Reilly Media, Inc. 1005 Gravenstein Highway North, Sebastopol, CA 95472 (707) 827-7000

------------------------------------------------------

.

------=_Part_-1790310585_179862840.1537371133215
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.=
w3.org/TR/html4/loose.dtd">
<html xmlns=3D"http://www.w3.org/1999/xhtml" xml:lang=3D"en" lang=3D"en">
<head>=20
<meta http-equiv=3D"Content-type" content=3D"text/html;charset=3DUTF-8">=20
<style type=3D"text/css">
@media only screen and (max-width:640px) {
.fullwidth {
  width:100% !important;
}
}

@media only screen and (max-width:475px) {
.mobile {
  width:100% !important;
  text-align:left !important;
  padding-top:15px !important;
  align:left !important;
}

.logos {
  padding-left:20px !important;
}
}
</style>=20
</head>=20
<body class=3D"fullwidth" style=3D"padding:0; margin:0; background-color:#f=
ff;"><style type=3D"text/css">div#emailPreHeader{ display: none !important;=
 }</style><div id=3D"emailPreHeader" style=3D"mso-hide:all; visibility:hidd=
en; opacity:0; color:transparent; mso-line-height-rule:exactly; line-height=
:0; font-size:0px; overflow:hidden; border-width:0; display:none !important=
;">Take the CKAD prep and exam at Velocity</div>=20
<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%" styl=
e=3D"font-family:'Open Sans',HelveticaNeue-Light,'Helvetica Neue Light',Hel=
vetica,sans-serif; background-color:#fff; font-size:16px; color:#444; line-=
height:22px; margin:0 0 12px 0; padding:30px 0 0 0;">=20
<tbody>
<tr>=20
<td bgcolor=3D"#ffffff">=20
<table border=3D"0" class=3D"fullwidth" cellpadding=3D"0" cellspacing=3D"0"=
 width=3D"640" align=3D"center">=20
<tbody>
<tr>=20
<td style=3D"padding:0 20px 30px 20px;">=20
<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"100%">=20
<tbody>
<tr>=20
<td style=3D"border-bottom:1px solid #ddd; padding-bottom:25px;">=20
<table align=3D"left" border=3D"0" cellpadding=3D"0" cellspacing=3D"0" styl=
e=3D"width:107px; height:30px; margin:0; padding:0 10px 0 0;">=20
<tbody>
<tr>=20
<td style=3D"width:107px; height:30px; padding:0; font-size:0; margin:0;"><=
a href=3D
"http://link.oreilly.com/P000Q0rFWM0rmy030d9S8U0" target=3D"_blank"
><img src=3D"https://cdn.oreillystatic.com/oreilly/email/logos/2018/oreilly=
-logo-107x30@2x.png" alt=3D"O'Reilly Media Logo" width=3D"107" height=3D"30=
" border=3D"0"></a> </td>=20
</tr>=20
</tbody>
</table>=20
<table align=3D"right" border=3D"0" cellpadding=3D"0" cellspacing=3D"0" wid=
th=3D"65%" class=3D"mobile" style=3D"width:70%; max-width:420px; height:30p=
x; vertical-align:top; padding:0; color:#666; font-family:'Open Sans',Helve=
ticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-weight:400=
; font-size:14px; line-height:14px; text-align:right;">=20
<tbody>
<tr>=20
<td><a href=3D
"http://link.oreilly.com/F8dMW00na0F0r03U000SrQy" target=3D"_blank" style=
=3D"font-weight:400; text-decoration:none; color:#666;"
>Learning Platform</a> =C2=B7 <a href=3D
"http://link.oreilly.com/JM0WFU0o8S0Qb3d0rr000y0" target=3D"_blank" style=
=3D"font-weight:400; text-decoration:none; color:#666;"
>Conferences</a> =C2=B7 <a href=3D
"http://link.oreilly.com/ic3F0QM8rpW00r000dU0S0y" target=3D"_blank" style=
=3D"font-weight:400; text-decoration:none; color:#666;"
>Ideas</a> </td>=20
</tr>=20
</tbody>
</table> </td>=20
</tr>=20
</tbody>
</table>=20
<!--/end logo -->=20
<!-- Begin: Body -->=20
<table cellpadding=3D"0" cellspacing=3D"0" border=3D"0" style=3D"font-size:=
16px; line-height:24px; font-weight:300; font-family:'Open Sans',HelveticaN=
eue-Light,'Helvetica Neue Light',Helvetica,sans-serif; padding:0; margin:0;=
">=20
<tbody>
<tr>=20
<td align=3D"left" style=3D"color:#444; font-family:'Open Sans',HelveticaNe=
ue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; line-=
height:24px; font-weight:300; padding:25px 0 25px 0; margin:0;"> <p align=
=3D"left" style=3D"color:#444; font-family:'Open Sans',HelveticaNeue-Light,=
'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; line-height:24=
px; font-weight:300; padding:0 0 25px 0; margin:0;">Hi Mathieu, </p> <p ali=
gn=3D"left" style=3D"color:#444; font-family:'Open Sans',HelveticaNeue-Ligh=
t,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; line-height:=
24px; font-weight:300; padding:0 0 25px 0; margin:0;">There's an important =
question that many hiring managers are asking software engineer candidates =
more often: </p> <p align=3D"left" style=3D"color:#444; font-family:'Open S=
ans',HelveticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-=
size:16px; line-height:24px; font-weight:300; padding:0 0 25px 0; margin:0;=
"><em>"Can you develop and maintain applications using Kubernetes?"</em> </=
p> <p align=3D"left" style=3D"color:#444; font-family:'Open Sans',Helvetica=
Neue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; lin=
e-height:24px; font-weight:300; padding:0 0 25px 0; margin:0;">What's more,=
 according to the <a href=3D
"http://link.oreilly.com/id3F0QM8rqW00r000dU0S0y" style=3D"color:#0000ee; t=
ext-decoration:none; font-weight:300;"
>2017 Linux Foundation and Dice Open Source Jobs Report</a>, <strong>half o=
f hiring managers prefer to hire candidates with certifications</strong>, 6=
5% seek expertise in Linux, and 70% seek expertise in cloud technologies. <=
/p> <p align=3D"left" style=3D"color:#444; font-family:'Open Sans',Helvetic=
aNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; li=
ne-height:24px; font-weight:300; padding:0 0 25px 0; margin:0;">That's wher=
e the <strong>Certified Kubernetes Application Developer (CKAD)</strong> pr=
ep course and exam comes in. This 2-day program at the O'Reilly Velocity Co=
nference is the perfect way to increase your marketability and value in the=
 job market, putting you ahead of the pack as this emerging platform grows =
in popularity. </p> <p align=3D"left" style=3D"color:#444; font-family:'Ope=
n Sans',HelveticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; fo=
nt-size:16px; line-height:24px; font-weight:300; padding:0 0 10px 0; margin=
:0;">The CKAD prep course and exam will give you: </p>=20
<ul align=3D"left" style=3D"color:#444; font-family:'Open Sans',HelveticaNe=
ue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; line-=
height:24px; font-weight:300; padding:0 0 25px 32px; margin:0;">=20
<li style=3D"padding:0 0 10px 0;">An overview of the key concepts and skill=
s you need to know to pass the official Certified Kubernetes Application De=
veloper (CKAD) exam </li>=20
<li style=3D"padding:0 0 10px 0;">A review of the core API primitives and b=
asic security concepts such as service accounts and Pod security policies <=
/li>=20
<li>An understanding of all Pod concepts and API schema, how to configure a=
pplications, and Pod monitoring and logging practices </li>=20
</ul> <p align=3D"left" style=3D"color:#444; font-family:'Open Sans',Helvet=
icaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; =
line-height:24px; font-weight:300; padding:0 0 25px 0; margin:0;">Most impo=
rtantly, you can leave Velocity a Certified Kubernetes Application Develope=
r. Just pick the date and location that works best for you: <strong>Septemb=
er 30=E2=80=93October 1 in New York City</strong>, or <strong>October 30=E2=
=80=9331 in London</strong>. </p>=20
<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" style=3D"width:100%=
;">=20
<tbody>
<tr>=20
<td align=3D"center" style=3D"color:#444; font-family:'Open Sans',Helvetica=
Neue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; lin=
e-height:24px; font-weight:300; padding:0 5px 25px 0; margin:0; width:45%;"=
><a href=3D
"http://link.oreilly.com/trSr0U03Fd0We00QMy0008r" style=3D"color:#0000ee; t=
ext-decoration:none; font-weight:300;"
>Register for Velocity in New&nbsp;York</a> </td>=20
<td align=3D"center" style=3D"color:#444; font-family:'Open Sans',Helvetica=
Neue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16px; lin=
e-height:24px; font-weight:300; padding:0 0 25px 5px; margin:0; width:45%;"=
><a href=3D
"http://link.oreilly.com/kF0fW0Q003r00MrUd80sy0S" style=3D"color:#0000ee; t=
ext-decoration:none; font-weight:300;"
>Register for Velocity in London</a> </td>=20
</tr>=20
</tbody>
</table> <p align=3D"left" style=3D"color:#444; font-family:'Open Sans',Hel=
veticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:16p=
x; line-height:24px; font-weight:300; padding:0 0 25px 0; margin:0;">See yo=
u at Velocity, </p> <p align=3D"left" style=3D"color:#444; font-family:'Ope=
n Sans',HelveticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-serif; fo=
nt-size:16px; line-height:24px; font-weight:300; padding:0; margin:0;">The =
O'Reilly Team </p> </td>=20
</tr>=20
</tbody>
</table>=20
<!-- /End:Body -->=20
<!--begin footer-->=20
<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" width=3D"600" bgcol=
or=3D"#ffffff" class=3D"fullwidth" style=3D"text-align:left; margin:0; padd=
ing:0;">=20
<tbody>
<tr>=20
<td style=3D"padding:30px 0 0 0; border-top:1px solid #ddd; color:#444; mar=
gin:0;">=20
<table border=3D"0" cellpadding=3D"0" cellspacing=3D"0" style=3D"margin:0; =
padding:0; height:30px;">=20
<tbody>
<tr>=20
<td style=3D"vertical-align:middle; padding:0;"><a href=3D
"http://link.oreilly.com/P000Q0rFWM0rmy030d9S8U0" target=3D"_blank"
><img src=3D"https://cdn.oreillystatic.com/oreilly/email/logos/2018/oreilly=
-logo-107x30@2x.png" alt=3D"O'Reilly Media Logo" width=3D"107" height=3D"30=
" border=3D"0"></a> </td>=20
<td class=3D"logos" style=3D"vertical-align:middle; padding:0 0 0 30px;"><a=
 href=3D
"http://link.oreilly.com/eQ0dtS00000MF30Uy8gWr0r" target=3D"_blank"
><img src=3D"https://cdn.oreillystatic.com/oreilly/email/logos/2018/twitter=
-20171204-30x24_360.png" alt=3D"Twitter Logo" width=3D"30" border=3D"0"></a=
> </td>=20
<td class=3D"logos" style=3D"vertical-align:middle; padding:0 0 0 30px;"><a=
 href=3D
"http://link.oreilly.com/p0y00F08hS0U0uQMrW0rd30" target=3D"_blank"
><img src=3D"https://cdn.oreillystatic.com/oreilly/email/logos/2018/faceboo=
k-20171204-14x24_360.png" alt=3D"Facebook Logo" width=3D"14" border=3D"0"><=
/a> </td>=20
<td class=3D"logos" style=3D"vertical-align:middle; padding:0 0 0 30px;"><a=
 href=3D
"http://link.oreilly.com/o00vQ00i3008rF0ydSMrUW0" target=3D"_blank"
><img src=3D"https://cdn.oreillystatic.com/oreilly/email/logos/2018/linkedi=
n-20171204-24x24_360.png" alt=3D"LinkedIn Logo" width=3D"24" border=3D"0"><=
/a> </td>=20
</tr>=20
</tbody>
</table> <p style=3D"font-family:'Open Sans',HelveticaNeue-Light,'Helvetica=
 Neue Light',Helvetica,sans-serif; font-size:14px; line-height:22px; font-w=
eight:300; color:#666; padding:0; margin:25px 0 0 0;">This message was sent=
 to mlecarme@bearstech.com. You are receiving this because you're a custome=
r of O'Reilly Media, or you've signed up to receive email from us. We hope =
you found this message to be useful. However, if you'd rather not receive f=
uture emails of this type from O'Reilly, please <a target=3D"_blank" href=
=3D
"http://link.oreilly.com/gw0dyQW0U008MS0rF000j3r" style=3D"font-weight:norm=
al; color:#666;"
>manage your preferences or unsubscribe here</a>. </p> <p style=3D"font-fam=
ily:'Open Sans',HelveticaNeue-Light,'Helvetica Neue Light',Helvetica,sans-s=
erif; font-size:14px; line-height:22px; font-weight:300; color:#666; paddin=
g:0; margin:15px 0 0 0;">Please read our <a target=3D"_blank" href=3D
"http://link.oreilly.com/ry0x0S030MkWdQF8r00U0r0" style=3D"font-weight:norm=
al; color:#666;"
>Privacy Policy</a>. </p> <p style=3D"font-family:'Open Sans',HelveticaNeue=
-Light,'Helvetica Neue Light',Helvetica,sans-serif; font-size:14px; line-he=
ight:22px; font-weight:300; color:#666; padding:0; margin:15px 0 0 0;">O'Re=
illy Media, Inc. 1005 Gravenstein Highway North, Sebastopol, CA 95472 (707)=
 827-7000 </p> </td>=20
</tr>=20
</tbody>
</table> </td>=20
</tr>=20
</tbody>
</table> </td>=20
</tr>=20
</tbody>
</table>=20
<!--/footer -->=20

<img src=3D"http://link.oreilly.com/trk?t=3D1&mid=3DMTA3LUZNUy0wNzA6MjU5Nzo=
2MjE5OjE5NDQ1OjA6NDA3Njo5OjgwMDc6ODMxOTI0OTptbGVjYXJtZUBiZWFyc3RlY2guY29t" =
width=3D"1" height=3D"1" style=3D"display:none !important;" alt=3D"" />

<!--This is a comment -->
</body>
</html>
------=_Part_-1790310585_179862840.1537371133215--
`
