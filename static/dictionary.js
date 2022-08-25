function getDictionary(notCache){
    if (notCache) {
        resetPaging();
    }
    showLoading();
    let url = "./api/dictionary";
    let params = { not_cache:false, level:"n5", start:0,page_size: 20,password:""};
    if (notCache) {
        params.not_cache = true;
        params.password = document.getElementsByName("sync-password")[0].value;
    }
    params.page_size = parseInt(document.getElementById("paging").getAttribute("data-page-size"),10)
    params.start = parseInt(document.getElementById("paging").getAttribute("data-start"),10)
    params.level = $('input[type=radio][name=oplevel]:checked').val()
    url = url + "?" + $.param(params);
    $.get(url)
        .done(function(data) {
            if(data.length > 0) {
                let content = "";
                if(params.start === 0){
                    $('#dict-table-body').html("");
                }
                data.forEach(function(row,index) {
                    let globalIdx = index + params.start
                    content = "<tr>";
                    content = content + "<td>" + (globalIdx + 1) + "</td>";
                    content = content + "<td>" + row.text + "</td>";
                    content = content + "<td>" + row.alphabet + "</td>";
                    content = content + "<td>" + row.mean_eng + "</td>";
                    content = content + "<td>" + row.mean_vn + "</td>";
                    let detail = "<td><div class='content-detail' id='content-detail-" + globalIdx + "'>" + row.detail + "</div><button class='btn btn-default btn-sm btn-detail' data-toggle='modal' data-target='.bd-example-modal-lg' data-detail='content-detail-" + globalIdx + "' onclick='showDetail(this)'><span class='glyphicon glyphicon-info-sign'></span> 詳細</button></td>";
                    if(row.detail.trim() === ""){
                        detail = "<td><div class='content-detail' id='content-detail-" + globalIdx + "'>" + row.detail + "</div><button class='btn btn-primary btn-sm btn-detail' data-detail='content-detail-" + globalIdx + "' onclick='getDetail(this,"+globalIdx+")'><span class='glyphicon glyphicon-cloud-download'></span> 取得</button></td>";
                    }
                    content = content + detail;
                    content = content + "</tr>";
                    $('#dict-table-body').append(content);
                });
                let pageSize = parseInt(document.getElementById("paging").getAttribute("data-page-size"),10)
                let prevStart = parseInt(document.getElementById("paging").getAttribute("data-start"),10)
                document.getElementById("paging").setAttribute("data-start",(prevStart+pageSize).toString(10))
                document.getElementById("paging").setAttribute("data-lock","false");
            }
            showPage();
        })
        .fail(function() {
            location.reload();
        })
}

window.onscroll = function(ev) {
    if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
        if(document.getElementById("paging").getAttribute("data-lock") === "false"){
            document.getElementById("paging").setAttribute("data-offset",document.body.offsetHeight.toString());
            getDictionary();
        }
    }
};

$('input[type=radio][name=oplevel]').change(function() {
    reload();
});

function reload(){
    resetPaging();
    getDictionary();
}

function resetPaging(){
    document.getElementById("paging").setAttribute("data-lock","false");
    document.getElementById("paging").setAttribute("data-start","0");
    document.getElementById("paging").setAttribute("data-offset","0");
}

function showDetail(el){
    document.getElementsByClassName("modal-content-detail")[0].innerHTML = "<div class='modal-detail'>" + document.getElementById($(el).attr("data-detail")).innerHTML + "</div>";
}

function getDetail(el,idx){
    el.disabled = "disabled";
    el.innerHTML = "Loading...";
    let url = "./api/dictionary/" + idx;
    let params = { level:"n5" };
    params.level = $('input[type=radio][name=oplevel]:checked').val()
    url = url + "?" + $.param(params);
    $.ajax({
        url: url,
        type: 'PUT',
        success: function (result) {
            document.getElementById($(el).attr("data-detail")).innerHTML = result;
            showDetail(el);
            el.outerHTML = "<button class='btn btn-default btn-sm btn-detail' data-toggle='modal' data-target='.bd-example-modal-lg' data-detail='content-detail-" + idx + "' onclick='showDetail(this)'><span class='glyphicon glyphicon-info-sign'></span> 詳細</button>";
            $('.bd-example-modal-lg').modal('toggle');
        },
    });
}

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

//Call getDictionary on page load.
getDictionary();