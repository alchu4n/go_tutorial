import request from '@/utils/request'

// 查询Article列表
export function listArticle(query) {
  return request({
    url: '/api/v1/article',
    method: 'get',
    params: query
  })
}

// 查询Article详细
export function getArticle(id) {
  return request({
    url: '/api/v1/article/' + id,
    method: 'get'
  })
}

// 新增Article
export function addArticle(data) {
  return request({
    url: '/api/v1/article',
    method: 'post',
    data: data
  })
}

// 修改Article
export function updateArticle(data) {
  return request({
    url: '/api/v1/article/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Article
export function delArticle(data) {
  return request({
    url: '/api/v1/article',
    method: 'delete',
    data: data
  })
}

