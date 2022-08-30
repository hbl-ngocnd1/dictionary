function getWonDerWord(notCache){
    if (notCache) {
        resetPaging();
    }
    showLoading();
    let url = "./api/wonder-word";

    $.get(url)
        .done(function(data) {
            if(data.length > 0) {
                let content = "";
                data.forEach(function(rows,idx1) {
                    rows.forEach(function(row,idx2) {
                        content = "<tr>";
                        content = content + "<td>" + (row.index + 1) + "</td>";
                        content = content + "<td>" + row.term + "</td>";
                        content = content + "<td>" + row.reading + "</td>";
                        content = content + "<td>" + row.explanation + "</td>";
                        content = content + "<td>" + row.mean + "</td>";
                        content = content + "<td>" + row.example.replace(/<[^>]*>?/gm, '') + "</td>";
                        content = content + "</tr>";
                        $('#dict-table-body').append(content);
                    });
                });
                // let pageSize = parseInt(document.getElementById("paging").getAttribute("data-page-size"),10)
                // let prevStart = parseInt(document.getElementById("paging").getAttribute("data-start"),10)
                // document.getElementById("paging").setAttribute("data-start",(prevStart+pageSize).toString(10))
                // document.getElementById("paging").setAttribute("data-lock","false");
            }
            showPage();
        })
        .fail(function() {
            location.reload();
        })
}

// window.onscroll = function(ev) {
//     if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
//         if(document.getElementById("paging").getAttribute("data-lock") === "false"){
//             document.getElementById("paging").setAttribute("data-offset",document.body.offsetHeight.toString());
//             getDictionary();
//         }
//     }
// };

// $('input[type=radio][name=oplevel]').change(function() {
//     reload();
// });

function reload(){
    resetPaging();
    getWonDerWord();
}

// function resetPaging(){
//     document.getElementById("paging").setAttribute("data-lock","false");
//     document.getElementById("paging").setAttribute("data-start","0");
//     document.getElementById("paging").setAttribute("data-offset","0");
// }

// function showDetail(el){
//     document.getElementsByClassName("modal-content-detail")[0].innerHTML = "<div class='modal-detail'>" + document.getElementById($(el).attr("data-detail")).innerHTML + "</div>";
// }

function showPage(){
    document.getElementById("loader").style.display = "none";
    document.getElementById("dict-table").style.display = "block";
    document.getElementById("paging").innerText = "";
    document.body.scrollTop = parseInt(document.getElementById("paging").getAttribute("data-offset"),10);
}

function showLoading(){
    document.getElementById("loader").style.display = "block";
    let start = parseInt(document.getElementById("paging").getAttribute("data-start"),10)
    if(start === 0){
        document.getElementById("dict-table").style.display = "none";
    }else{
        document.getElementById("paging").innerText = "loading more...";
    }
    document.getElementById("paging").setAttribute("data-lock","true");
}

//Call getWonDerWord on page load.
getWonDerWord();