$("#wda-upload-form").submit(function (e) {
    e.preventDefault();
    var formData = new FormData();
    formData.append('file', $('#wda-input-file')[0].files[0]);

    $.ajax({
        url: '/upload-wda',
        type: 'POST',
        data: formData,
        processData: false,  // tell jQuery not to process the data
        contentType: false,  // tell jQuery not to set contentType
        success: function (data) {
            console.log(data);
            alert(data);
            location.reload()
        },
        error: function (data) {
            console.log(data);
            alert(data);
            location.reload()
        }
    });
});

/* Show info modal with provided text */
function showConfigModal() {
    /* Get the modal element */
    var modal = document.getElementById("configuration-modal")

    /* Get the close button */
    var span = document.getElementsByClassName("close")[0]

    $("#modal-body").load("/static/project_config_submission_form.html");

    /* Display the modal blocking interaction */
    modal.style.display = "block";

    /* Close the modal if you click anywhere outside the modal */
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
};

function showDockerfileModal() {
    /* Get the modal element */
    var modal = document.getElementById("configuration-modal")

    /* Get the close button */
    var span = document.getElementsByClassName("close")[0]

    $("#modal-body").load("/static/edit_file_modal.html");

    /* Display the modal blocking interaction */
    modal.style.display = "block";

    /* Close the modal if you click anywhere outside the modal */
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
}

function buildImage(obj) {
    alert("Building Docker image, this could take a while.")
    var test = document.getElementById("image-status-cell")
    test.textContent = "Image building"
    /* Call the endpoint that will restart the selected container */
    var req = $.ajax({
        async: true,
        type: "GET",
        url: "/build-image"
    });
}

function removeImage() {
    /* Show loading indicator until response is returned */
    $('#loading').css("visibility", "visible");

    /* Call the endpoint that will restart the selected container */
    $.ajax({
        dataType: false,
        async: true,
        type: "GET",
        url: "/remove-image",
        success: function (data) {
            alert(data)
        },
        error: function (data) {
            alert(JSON.stringify(data))
        }
    });

    /* Hide the loading after response is returned */
    $('#loading').css("visibility", "hidden");
}

function removeUdevListener() {
    var sudoPasswordStatus = document.getElementById("sudo-password-cell").getAttribute("value")
    if (sudoPasswordStatus === "false") {
        alert("Elevated permissions are needed to perform this action. Please set your user sudo password in the '.config.yaml' file.")
        return
    }
    /* Show loading indicator until response is returned */
    $('#loading').css("visibility", "visible");

    /* Call the endpoint that will restart the selected container */
    $.ajax({
        dataType: false,
        async: true,
        type: "GET",
        url: "/remove-udev-listener",
        success: function (data) {
            alert(data)
            /* Reload the page to get the new info */
            location.reload();
        },
        error: function (data) {
            alert(JSON.stringify(data))
            /* Reload the page to get the new info */
            location.reload();
        }
    });

    /* Hide the loading after response is returned */
    $('#loading').css("visibility", "hidden");
}

function setupUdevListener(gridBool) {
    var sudoPasswordStatus = document.getElementById("sudo-password-cell").getAttribute("value")
    if (sudoPasswordStatus === "false") {
        alert("Elevated permissions are needed to perform this action. Please set your user sudo password in the '.config.yaml' file.")
        return
    }
    var imageStatus = document.getElementById("image-status-cell").getAttribute("value")
    if (imageStatus === "is-not-available") {
        alert("The 'ios-appium' Docker image is not available, listener will not be started.")
        return
    }
    /* Show loading indicator until response is returned */
    $('#loading').css("visibility", "visible");

    if (gridBool) {
        url = "/setup-udev-listener"
    } else {
        url = "/setup-udev-listener"
    }

    /* Call the endpoint that will start the respective listener config */
    $.ajax({
        dataType: false,
        async: true,
        type: "GET",
        url: url,
        success: function (data) {
            alert(data)
            /* Reload the page to get the new info */
            location.reload();
        },
        error: function (data) {
            alert(JSON.stringify(data))
            /* Reload the page to get the new info */
            location.reload();
        }
    });

    /* Hide the loading after response is returned */
    $('#loading').css("visibility", "hidden");
}

function showIOSDeviceSelection() {
    $.ajax({
        contentType: 'application/json',
        async: true,
        type: "GET",
        url: "/ios-devices",
        success: function (data) {
            $('#add-device-button').prop('disabled', true);
            /* Get the modal element */
            var modal = document.getElementById("device-selection-modal")

            /* Get the close button */
            var span = document.getElementsByClassName("close")[0]

            let dropdown = $('#device-dropdown');

            dropdown.empty();

            dropdown.append('<option>Choose device</option>');
            dropdown.prop('selectedIndex', 0);
            var response = JSON.parse(data)
            for (let i = 0; i < response.deviceList.length; i++) {
                dropdown.append($('<option></option>').attr('value', response.deviceList[i].Udid + ":" + response.deviceList[i].ProductVersion).text("Device UDID: " + response.deviceList[i].Udid + ", Product Type: " + response.deviceList[i].ProductType + ", Device OS: " + response.deviceList[i].ProductVersion));
            }
            // Clear the previous value in the device name input if any
            $("#register-device-name").val("");

            /* Display the modal blocking interaction */
            modal.style.display = "block";

            /* Close the modal if you click anywhere outside the modal */
            window.onclick = function (event) {
                if (event.target == modal) {
                    modal.style.display = "none";
                }
            }
        },
        error: function (data) {
            alert($.parseJSON(data.responseText).error_message)
        }
    });
}

function registerIOSDevice() {
    var modal = document.getElementById("device-selection-modal")
    // Get the device UDID from the value of the selected option
    var device_info = $("#device-dropdown").val();
    var device_name = $("#register-device-name").val();
    if (device_name === "") {
        alert("Please provide a device name. Avoid using special characters and spaces except for '_'. Example: iPhone_11")
        return
    }
    deviceInfoArray = device_info.split(new RegExp(":"));

    // Send a request to register the device with the respective selected device UDID
    $.ajax({
        contentType: 'application/json',
        async: true,
        type: "POST",
        data: JSON.stringify({ "device_udid": deviceInfoArray[0], "device_os_version": deviceInfoArray[1] }),
        url: "/ios-devices/register",
        success: function (data) {
            alert("Successfully added device with UDID: " + deviceInfoArray[0] + " to the config.json file.")
            modal.style.display = "none";
        },
        error: function (data) {
            alert($.parseJSON(data.responseText).error_message)
            modal.style.display = "none";
        }
    });
}

function showSudoPasswordInput() {

    var modal = document.getElementById("sudo-password-input-modal")

    /* Display the modal blocking interaction */
    modal.style.display = "block";

    /* Close the modal if you click anywhere outside the modal */
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
}

function setSudoPassword() {
    var modal = document.getElementById("sudo-password-input-modal")
    var sudo_password = $("#sudo-password-input").val();
    if (sudo_password === "") {
        alert("Please provide a non-empty value for the sudo password.")
        return
    }
    // Send a request to register the device with the respective selected device UDID
    $.ajax({
        contentType: 'application/json',
        async: true,
        type: "POST",
        data: JSON.stringify({ "sudo_password": sudo_password }),
        url: "/set-sudo-password",
        success: function (data) {
            alert("Successfully set sudo password")
            modal.style.display = "none";
            location.reload()
        },
        error: function (data) {
            alert("Couldn't set sudo password")
            modal.style.display = "none";
            location.reload()
        }
    });
}

function showWDAUploadForm() {
    var modal = document.getElementById("upload-wda-modal")

    /* Display the modal blocking interaction */
    modal.style.display = "block";

    /* Close the modal if you click anywhere outside the modal */
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
}

function enableDisableSubmit(dropdownObj) {
    // Check the currently selected option
    // If it is the default one disable the Add button else enable it
    if (dropdownObj.options[dropdownObj.selectedIndex].text === "Choose device") {
        $('#add-device-button').prop('disabled', true);
    } else {
        $('#add-device-button').prop('disabled', false);
    }
}

function notImplemented() {
    alert("Not implemented")
}

window.addEventListener("DOMContentLoaded", function () {
    var socket = new WebSocket("ws://localhost:10000/ws");
    socket.onmessage = function (e) {
        alert(e.data)
        location.reload()
    };
    setWDAStatusCellBackground()
    setImageStatusCellBackground()
    setUdevIOSListenerCellBackground()
    setSudoPasswordCellBackground()
});

function setWDAStatusCellBackground() {
    let statusCell = document.getElementById("wda-status-cell")
    if (statusCell.textContent === "true") {
        statusCell.style.backgroundColor = "#4CAF50"
        statusCell.setAttribute("value", "true")
    } else {
        statusCell.style.backgroundColor = "#eb4f34"
        statusCell.setAttribute("value", "false")
    }
}

function setImageStatusCellBackground() {
    let statusCell = document.getElementById("image-status-cell");
    if (statusCell.textContent === "Image available") {
        statusCell.style.backgroundColor = "#4CAF50"
        statusCell.setAttribute("value", "is-available")
    } else if (statusCell.textContent === "Image not available") {
        statusCell.style.backgroundColor = "#eb4f34"
        statusCell.setAttribute("value", "is-not-available")
    } else {
        statusCell.style.backgroundColor = "#ebe134"
        statusCell.setAttribute("value", "undefined")
    }
}

function setUdevIOSListenerCellBackground() {
    let statusCell = document.getElementById("udev-ios-listener-status-cell");
    if (statusCell.textContent === "Udev rules set.") {
        statusCell.style.backgroundColor = "#4CAF50"
        statusCell.setAttribute("value", "is-running")
    } else if (statusCell.textContent === "Udev rules not set.") {
        statusCell.style.backgroundColor = "#eb4f34"
        statusCell.setAttribute("value", "is-not-running")
    }
}

function setSudoPasswordCellBackground() {
    let statusCell = document.getElementById("sudo-password-cell")
    if (statusCell.textContent === "true") {
        statusCell.style.backgroundColor = "#4CAF50"
        statusCell.setAttribute("value", "true")
    } else {
        statusCell.style.backgroundColor = "#eb4f34"
        statusCell.setAttribute("value", "false")
    }
}