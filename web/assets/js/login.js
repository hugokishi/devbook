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
    window.location = "/home"
  }).fail(function() {
    alert("Usuário ou senha inválidos!");
  }); 
}