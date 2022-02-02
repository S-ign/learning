function SweetAlert() {
	let Msg = function (c) {
		const {
			toast = true,
			msg = "",
			footerMsg = "",
			icon = "", //[success,error,warning,info,quesion]
			position = "", //top, bottom, mid, start, end, top-end, bottom-start
			timer = 3000,
		} = c;

		const Toast = Swal.mixin({
			toast: toast,
			icon: icon,
			title: msg,
			position: position,
			footer: footerMsg,
			showConfirmButton: false,
			timer: timer,
			timerProgressBar: true,
			didOpen: (toast) => {
				toast.addEventListener('mouseenter', Swal.stopTimer)
				toast.addEventListener('mouseleave', Swal.resumeTimer)
			}
		})

		Toast.fire({})
	}

	let Form = async function (c) {
		const {
			title = "",
			html = "",
		} = c;
		const { value: formValues } = await Swal.fire({
			title: title,
			html: html,
			focusConfirm: false,
			showCancelButton: true,
			willOpen: () => {
				const elem = document.getElementById('reservation-dates-modal')
				const rangePicker = new DateRangePicker(elem, {
					showOnFocus: false
				})
			},
			preConfirm: () => {
				return [
					document.getElementById('start').value,
					document.getElementById('end').value
				]
			},
			didOpen: () => {
				let datepickerElements = document.getElementsByClassName('datepicker')
				for ( var i = 0; i < datepickerElements.length; i++) {
					datepickerElements[i].style.zIndex = "100000";
				}
			}
		})

		if (formValues) {
			if (formValues.dismiss !== Swal.DismissReason.cancel){
				if (formValues.value !== "") {
					if (c.callback !== undefined) {
						c.callback(formValues);
					} else {
						c.callback(false);
					}
				} else {
					c.callback(false);
				}
			}
		}

	}

	return {
		Msg: Msg,
		Form: Form,
	}
}

sweetAlert = SweetAlert()

let reservationNow = document.getElementsByClassName('reservation-now-datepicker')
if (reservationNow.length > 0) {
	for ( var i = 0; i < reservationNow.length; i++) {
			reservationNow[i].onclick = () => sweetAlert.Form({title:"Find Reservation",
			html: `
			<form id="check-availability-form" class="needs-validation" action="/reservation-date" novalidate method="post">
				<div class="row">
					<div class="col">
						<div id="reservation-dates-modal" class="row">
							<div class="col">
								<input class="form-control" type="text" name="start" id="start" placeholder="Arrival">
							</div>
							<div class="col">
								<input class="form-control" type="text" name="end" id="end" placeholder="Departure">
							</div>
						</div>
					</div>
				</div>
			</form>
			`,
				callback: function(formValues) {
					console.log("called");

					let tkn = document.querySelector('meta[name="csrf-token"]').content;
					let form = document.getElementById('check-availability-form');
					let formData = new FormData(form);
					formData.append("csrf_token", tkn);

					fetch('/search-availability-json', {
						method: 'POST',
						body: formData,
					})

						.then(response => response.json())
						.then(data => {
							console.log(data.ok);
							console.log(data.message);
						})
				}

			})
	}
}
