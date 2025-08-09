$(function() {
    // 从HTML元素的data属性中获取参数
    const paginationElement = $('#pagination');
    const totalPages = parseInt(paginationElement.data('total-pages'));
    const currentPage = parseInt(paginationElement.data('current-page'));
    const categoryId = paginationElement.data('category-id');

    // 初始化分页器
    paginationElement.jqPaginator({
        totalPages: totalPages,
        visiblePages: 10,
        currentPage: currentPage,
        onPageChange: function(num, type) {
            if (type == "change") {
                location.href = "/category" + categoryId + "?page=" + num;
            }
        }
    });
});
