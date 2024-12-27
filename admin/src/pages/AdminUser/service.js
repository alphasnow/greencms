import { request } from '@umijs/max';
import qs from 'qs';

export function getList(params) {
  return request('/api/admin/admin-user/index', {
    method: 'GET',
    params: params,
    paramsSerializer: function (params) {
      return qs.stringify(params, { arrayFormat: 'brackets' });
    },
  });
}

export function postStore(data) {
  return request('/api/admin/admin-user/create', {
    method: 'POST',
    data: data,
  });
}

export function getData(id) {
  return request('/api/admin/admin-user/show/' + id, {
    method: 'GET',
  });
}

export function postUpdate(id, data) {
  return request('/api/admin/admin-user/edit/' + id, {
    method: 'POST',
    data: data,
  });
}

export function postDelete(id) {
  return request('/api/admin/admin-user/delete/' + id, {
    method: 'POST',
  });
}

export function getOptions() {
  return request('/api/admin/admin-user/options', {
    method: 'GET',
  });
}