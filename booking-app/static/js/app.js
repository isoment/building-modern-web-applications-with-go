// Define a Prompt module we want to be able to use throughout our application to handle 
// alerts, popups etc.
function Prompt() {
  const notify = (msg, type) => {
    notie.alert({
      type: type,
      text: msg,
    })
  }

  const success = (c) => {
    const {
      msg = "",
      title = "",
      footer = "",
    } = c;

    Swal.fire({
      icon: 'success',
      title: title,
      text: msg,
      footer: footer,
    });
  }

  const error = (c) => {
    const {
      msg = "",
      title = "",
      footer = "",
    } = c;

    Swal.fire({
      icon: 'error',
      title: title,
      text: msg,
      footer: footer,
    });
  }

  const toast = (c) => {
    const {
      msg = '',
      icon = 'success',
      position = 'top-end'
    } = c;

    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      position: position,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      }
    });

    Toast.fire({});
  }

  const custom = async (c) => {
    const { icon = "", msg = "", title = "", showConfirmButton = true } = c;

    const { value: result } = await Swal.fire({
      icon: icon,
      title: title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton: showConfirmButton,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      didOpen: () => {
        if (c.didOpen !== undefined) {
          c.didOpen();
        }
      },
      preConfirm: () => {
        return [
          document.getElementById('start').value,
          document.getElementById('end').value
        ]
      }
    });

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  }

  return {
    notify,
    success,
    toast,
    error,
    custom
  }
}

function RoomAvailability(roomId, csrfToken) {
  document.getElementById("check-availability-button").addEventListener("click", function () {
    // Define the html form
    let html = `
    <form id="check-availability-form" action="" method="POST" novalidate class="needs-validation">
      <div class="row">
        <div class="col">
          <div id="reservation-dates-modal" class="row">
            <div class="col">
              <input disabled class="form-control" type="text" name="start" id="start" required placeholder="Arrival">
            </div>
            <div class="col">
              <input disabled class="form-control" type="text" name="end" id="end" required placeholder="Departure">
            </div>
          </div>
        </div>
      </div>
    </form>
    `;

    // Open a sweet alert passing the callback into it using fetch() to make a request.
    attention.custom({
      msg: html, 
      title: "Choose your dates",

      willOpen: () => {
        const elem = document.getElementById('reservation-dates-modal');
        const rp = new DateRangePicker(elem, {
          format: 'yyyy-mm-dd',
          showOnFocus: true,
          minDate: new Date(),
        })
      },

      didOpen: () => {
        document.getElementById('start').removeAttribute('disabled');
        document.getElementById('end').removeAttribute('disabled');
      },

      callback: function(result) {
        let form = document.getElementById("check-availability-form");
        let formData = new FormData(form);
        formData.append("csrf_token", csrfToken)
        formData.append("room_id", roomId);

        fetch('/search-availability-json', {
          method: "post",
          body: formData,
        })
          .then(response => {
            return response.json()
          })
          .then(data => {
            if (data.ok) {
              attention.custom({
                icon: 'success',
                msg: '<p>Room is available!</p>'
                  + '<p><a href="/book-room?id='
                  + data.room_id
                  + '&s='
                  + data.start_date
                  + '&e='
                  + data.end_date
                  + '" class="btn btn-primary">'
                  + 'Book now!</a></p>',
                showConfirmButton: false,
              })
            } else {
              attention.notify("No Availability", "error")
            }
          })
      }
    });
  });
}