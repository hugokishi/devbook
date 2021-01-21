$('#follow').on('click', follow)
$('#unfollow').on('click', unfollow)
$('#edit-user').on('submit', editUser)
$('#update-password').on('submit', updatePassword)
$('#delete-user').on('click', deleteUser)

function follow() {
  const userId = $(this).data('user-id')

  $(this).prop('disabled', true)

  $.ajax({
    url: `/users/${userId}/follow`,
    method: 'POST'
  }).done(function () {
    window.location = `/users/${userId}`
  }).fail(function () {
    Swal.fire(
      "Ops...",
      "Erro ao seguir o usuário!",
      "error"
    )
    $('#follow').prop('disabled', false)
  })
}

function unfollow() {
  const userId = $(this).data('user-id')

  $(this).prop('disabled', true)

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: 'POST'
  }).done(function () {
    window.location = `/users/${userId}`
  }).fail(function () {
    Swal.fire(
      "Ops...",
      "Erro ao parar de seguir o usuário!",
      "error"
    )
    $('#unfollow').prop('disabled', false)
  })
}

function editUser(e) {
  e.preventDefault()

  $.ajax({
    url: '/edit-profile',
    method: 'PUT',
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
    }
  }).done(function () {
    Swal.fire(
      "Sucesso!",
      "Usuário editado com sucesso!",
      "success"
    ).then(function () {
      window.location = "/profile"
    })
  }).fail(function () {
    Swal.fire(
      "Ops...",
      "Erro ao atualizar o usuário!",
      "error"
    )
  })
}

function updatePassword(e) {
  e.preventDefault()

  if ($('#new-pass').val() != $('#confirm-pass').val()) {
    Swal.fire(
      "Ops...",
      "As senhas não coinciedem!",
      "warning"
    )
    return;
  }

  $.ajax({
    url: '/edit-password',
    method: 'POST',
    data: {
      now: $('#now-pass').val(),
      new: $('#new-pass').val(),
    }
  }).done(function () {
    Swal.fire(
      "Sucesso!",
      "Senha editada com sucesso!",
      "success"
    ).then(function () {
      window.location = "/profile"
    })
  }).fail(function () {
    Swal.fire(
      "Ops...",
      "Erro ao atualizar a senha do usuário!",
      "error"
    )
  })
}

function deleteUser() {
  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning"
  }).then(function (confirm) {
    if (confirm.value) {
      $.ajax({
        url: "/delete-user",
        method: "DELETE"
      }).done(function () {
        Swal.fire("Sucesso!", "Seu usuário foi excluído com sucesso!", "success")
          .then(function () {
            window.location = "/logout";
          })
      }).fail(function () {
        Swal.fire("Ops...", "Ocorreu um erro ao excluir o seu usuário!", "error");
      });
    }
  })
}