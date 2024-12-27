import { request } from '@umijs/max';
import qs from 'qs';

export function getList(params) {
  return request('/api/admin/article-category/index', {
    method: 'GET',
    params: params,
    paramsSerializer: function (params) {
      return qs.stringify(params, { arrayFormat: 'brackets' });
    },
  });
}

export function postStore(data) {
  return request('/api/admin/article-category/create', {
    method: 'POST',
    data: data,
  });
}

export function getData(id) {
  return request('/api/admin/article-category/show/' + id, {
    method: 'GET',
  });
}

export function postUpdate(id, data) {
  return request('/api/admin/article-category/edit/' + id, {
    method: 'POST',
    data: data,
  });
}

export function postDelete(id) {
  return request('/api/admin/article-category/delete/' + id, {
    method: 'POST',
  });
}
