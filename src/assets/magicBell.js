function initializeNotifications(target, userEmail) {
    (function(i,s,o,g,r,a,m) {i['MagicBellObject'] = r;(i[r]=i[r] || function() {

        (i[r].q = i[r].q || []).push(arguments);}),(i[r].l = 1 * new Date());(a = s.createElement(o)), (
    
        m = s.getElementsByTagName(o)[0]);a.async = 1;a.src = g;m.parentNode.insertBefore(a, m);
    
      })(window,document,'script','https://unpkg.com/@magicbell/embeddable/dist/magicbell.min.js','magicbell');
      var target = document.getElementById("magicbell-inbox");
      var options = {
        apiKey: "57a892b7f9e5e9319fbda8fb1abef85a01d1652f",
        userEmail: "erob7856@gmail.com", // Replace with the logged-in user's email
        height: 500,
        theme: {"icon":{"borderColor":"#EDEDEF","width":"24px"},"banner":{"fontSize":"14px","backgroundColor":"#F8F5FF","textColor":"#3A424D","backgroundOpacity":1},"unseenBadge":{"backgroundColor":"#F80808"},"header":{"fontFamily":"inherit","fontSize":"15px","backgroundColor":"#FFFFFF","textColor":"#5225C1","borderRadius":"16px"},"footer":{"fontSize":"15px","backgroundColor":"#FFFFFF","textColor":"#5225C1","borderRadius":"16px"},"notification":{"default":{"fontFamily":"inherit","fontSize":"14px","textColor":"#3A424D","borderRadius":"16px","backgroundColor":"#FFFFFF","hover":{"backgroundColor":"#F2EDFC"},"state":{"color":"transparent"},"margin":"8px"},"unseen":{"textColor":"#3A424D","backgroundColor":"#F8F5FF","hover":{"backgroundColor":"#F2EDFC"},"state":{"color":"#5225C1"}},"unread":{"textColor":"#3A424D","backgroundColor":"#F8F5FF","hover":{"backgroundColor":"#F2EDFC"},"state":{"color":"#5225C1"}}}},
        locale: "en",
      };
    
      magicbell('render', target, options);
}