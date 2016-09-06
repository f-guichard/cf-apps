(function (){

   var infos = { description : 'adresse mail eui', mail : 'v2@orange.fr' };

   var euimod = angular.module('module_eui', []);

   euimod.controller('euiController', [ '$http', '$log', function($http, $log) {
       this.infos = infos;
       
       this.validateForm = function() {
         $log.warn('Valide !');
         $http.get('http://cors.exmaple.com/authen?login=&pass=')
                  .success(function (rauhen) {
                            $log.warn('Recu ' + rauthen);
                          });
       };
   }]);
})();

