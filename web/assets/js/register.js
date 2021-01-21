$('#form-register').on('submit', createUser)

function createUser(e){
  e.preventDefault();

  if($('#pass').val() != $('#pass-confirm').val()){
    Swal.fire('Ops...', "As senhas não coincidem!", "error")
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      password: $('#pass').val()
    }
  }).done(function(){
    Swal.fire('Sucesso', "Usuário cadastrado com sucesso!", "success").then(function(){
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $("#email").val(),
          password: $('#pass').val(),
        }
      }).done(function(){
        window.location = "/feed"
      }).fail(function(){
        Swal.fire('Ops...', "Erro ao autenticar o usuário!", "error")
      })
    })
  }).fail(function(err) {
    Swal.fire('Ops...', "Erro ao cadastrar o usuário!", "error")
  }); 
}