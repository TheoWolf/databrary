'use strict';

app.directive('userPasswordForm', [
  'messageService', 'constantsService', 'modelService',
  function (messages, constants, models) {
    var link = function ($scope) {
      var form = $scope.userPasswordForm;

      var init = function () {
        form.data = {
          email: undefined,
        };
      };

      //

      form.resetSuccessFn = undefined;

      form.send = function () {
        messages.clear(form);
        models.Login.issuePassword($scope.userPasswordForm.data)
          .then(function () {
            form.validator.server({});

            messages.add({
              type: 'green',
              body: constants.message('reset.request.success', form.data.email),
              owner: form
            });

            init();
          }, function (res) {
            form.validator.server(res);
          });
      };

      form.validator.client({
        email: {
          tips: constants.message('reset.email.help'),
          errors: constants.message('login.email.error'),
        },
      }, true);
    };

    //

    return {
      restrict: 'E',
      templateUrl: 'party/userPasswordForm.html',
      scope: false,
      replace: true,
      link: link
    };
  }
]);
