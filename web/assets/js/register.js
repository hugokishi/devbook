$('#form-register').on('submit', createUser)

function createUser(e){
  e.preventDefault();

  if($('#pass').val() != $('#pass-confirm').val()){
    alert("As senhas estão diferentes!");
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
    alert("Usuário cadastrado com sucesso!");
  }).fail(function(err) {
    console.log(err)
    alert("Não foi possivel cadastrar o usuário");
  }); 
}