(function(){
	'use strict';

	angular
		.module('shipyard.registry')
		.controller('RegistriesController', RegistriesController);

	RegistriesController.$inject = ['resolvedRegistries', 'RegistryService', '$state', '$timeout', '$http'];
	function RegistriesController(resolvedRegistries, RegistryService, $state, $timeout, $http) {
            var vm = this;
            vm.registries = resolvedRegistries;
            vm.refresh = refresh;
            vm.registriesInUse = [];
            vm.selectedRegistry = "";
            vm.showRemoveRegistryDialog = showRemoveRegistryDialog;
            vm.removeRegistry = removeRegistry;
            vm.logRegistries = logRegistries;
            vm.loginRegistry = loginRegistry;

            function refresh() {
                RegistryService.list()
                    .then(function(data) {
                        vm.registries = data; 
                    }, function(data) {
                        vm.error = data;
                    });
                vm.error = "";
            }

            function showRemoveRegistryDialog(registry) {
                vm.selectedRegistry = registry;
                $('.ui.small.remove.modal').modal('show');
            }

            function removeRegistry() {
                RegistryService.removeRegistry(vm.selectedRegistry)
                    .then(function(data) {
                        vm.refresh();
                    }, function(data) {
                        vm.error = data;
                    });
            }

            function logRegistries(index) {
                console.log(vm.registries[index]);
                console.log(index);
            }

            function isValid() {
                // TODO: health check registries
                return true;
            }

            function loginRegistry(index) {
                if (!isValid()) {
                    return;
                }
                vm.request = vm.registries[index];
                $http
                    .post('/api/docker/login', vm.request)
                    .success(function(data, status, headers, config) {
                        $state.transitionTo('dashboard.registry');
                    })
                    .error(function(data, status, headers, config) {
                        vm.error = data;
                    });
            }
	}
})();
