<template>
  <div>
    <h1>📦 庫存管理</h1>
    <p>管理飼料、藥品和其他養殖材料的庫存。</p>

    <div class="card">
      <div class="card-header">
        <h5>庫存清單</h5>
      </div>
      <div class="card-body">
        <table class="table">
          <thead>
            <tr>
              <th>物品名稱</th>
              <th>數量</th>
              <th>單位</th>
              <th>最後更新</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>飼料 A</td>
              <td>500</td>
              <td>kg</td>
              <td>2026-04-07</td>
            </tr>
            <tr>
              <td>藥品 B</td>
              <td>20</td>
              <td>瓶</td>
              <td>2026-04-06</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Assets'
}
</script>