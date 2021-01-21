$('#form-login').on('submit', authenticateUser)

function authenticateUser(e){
  e.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      password: $('#pass').val()
    }
  }).done(function(){
    window.location = "/feed"
  }).fail(function() {
    Swal.fire('Ops...', "Usuário ou senha inválidos", "error")
  }); 
}