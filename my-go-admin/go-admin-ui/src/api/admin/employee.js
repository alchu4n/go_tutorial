import request from '@/utils/request'

// 查询Employee列表
export function listEmployee(query) {
  return request({
    url: '/api/v1/employee',
    method: 'get',
    params: query
  })
}

// 查询Employee详细
export function getEmployee(id) {
  return request({
    url: '/api/v1/employee/' + id,
    method: 'get'
  })
}

// 新增Employee
export function addEmployee(data) {
  return request({
    url: '/api/v1/employee',
    method: 'post',
    data: data
  })
}

// 修改Employee
export function updateEmployee(data) {
  return request({
    url: '/api/v1/employee/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Employee
export function delEmployee(data) {
  return request({
    url: '/api/v1/employee',
    method: 'delete',
    data: data
  })
}

