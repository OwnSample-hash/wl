function ValidateField(id, validator, errorMessage) {
  if (validator(id)) {
    $(`#${id}`).removeClass("is-invalid");
    return true;
  } else {
    $(`#${id}`).addClass("is-invalid");
    $(`#${id}`).attr("placeholder", errorMessage);
    return false;
  }
}

function NumberValidator(id) {
  const value = $(`#${id}`).val();
  return !isNaN(value) && value > 0;
}

function EmailValidator(id) {
  const value = $(`#${id}`).val();
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return regex.test(value);
}

function PhoneValidator(id) {
  const value = $(`#${id}`).val();
  const regex = /^\d{10}$/; // Example: 10-digit phone number
  return regex.test(value);
}

function StringValidator(id) {
  const value = $(`#${id}`).val();
  return value && value.trim() !== "";
}

const ids = [
  "firstName",
  "lastName",
  "email",
  "phone",
  "address",
  "city",
  "state",
  "zip",
  "country",
];

function ValidateAll() {
  let isValid = true;
  ids.forEach((id) => {
    switch (id) {
      case "firstName":
        ValidateField(id, StringValidator, "Please enter a valid name");
        break;
      case "lastName":
        ValidateField(id, StringValidator, "Please enter a valid name");
        break;
      case "email":
        ValidateField(id, EmailValidator, "Please enter a valid email");
        break;
      case "phone":
        ValidateField(id, PhoneValidator, "Please enter a valid phone number");
        break;
      case "address":
        ValidateField(id, StringValidator, "Please enter a valid address");
        break;
      case "city":
        ValidateField(id, StringValidator, "Please enter a valid city");
        break;
      case "state":
        ValidateField(id, StringValidator, "Please enter a valid state");
        break;
      case "zip":
        ValidateField(id, NumberValidator, "Please enter a valid zip code");
        break;
      case "country":
        ValidateField(id, StringValidator, "Please enter a valid country");
        break;
    }
  });
  return isValid;
}

function submitForm(id, csrfToken) {
  if (!ValidateAll()) {
    return;
  }
  // Perform form submission
}
