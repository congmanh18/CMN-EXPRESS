
### **1. Các trạng thái chính**
| **Trạng thái (Tiếng Việt)**     | **Trạng thái (Tiếng Anh)**         | **Giải nghĩa**                                                                 |
|----------------------------------|-------------------------------------|--------------------------------------------------------------------------------|
| Đã lấy hàng                     | `Picked up`                        | Shipper đã nhận hàng từ người gửi.                                            |
| Đang vận chuyển                 | `In transit`                       | Hàng hóa đang được vận chuyển đến người nhận.                                |
| Đang giao hàng                  | `Out for delivery`                 | Shipper đang giao hàng đến địa chỉ của người nhận.                           |
| Chờ phát lại                    | `Awaiting re-delivery`             | Hàng hóa không giao được và đang chờ giao lại.                               |
| Giao thành công                | `Delivered successfully`           | Hàng hóa đã được giao đến người nhận thành công.                             |
| Chờ xử lý                      | `Pending processing`               | Đơn hàng đang chờ xử lý bởi hệ thống hoặc nhân viên.                         |
| Đã duyệt hoàn                  | `Return approved`                  | Yêu cầu hoàn trả hàng đã được phê duyệt.                                     |
| Đang chuyển hoàn               | `Returning to sender`              | Hàng hóa đang được vận chuyển ngược lại cho người gửi.                       |
| Đã trả                         | `Returned`                         | Hàng hóa đã được trả lại cho người gửi.                                      |
| Phát tiếp                      | `Re-delivering`                    | Hàng hóa đang được giao lại lần tiếp theo.                                   |

---

### **2. Các trạng thái khác**
| **Trạng thái (Tiếng Việt)**     | **Trạng thái (Tiếng Anh)**         | **Giải nghĩa**                                                                 |
|----------------------------------|-------------------------------------|--------------------------------------------------------------------------------|
| Tạo mới                        | `Created`                          | Đơn hàng mới được tạo và lưu vào hệ thống.                                    |
| Đã tiếp nhận                   | `Received`                         | Đơn hàng đã được tiếp nhận trong hệ thống.                                    |
| Đang lấy hàng                  | `Picking up`                       | Shipper đang trên đường đến nhận hàng từ người gửi.                           |
| Tồn - Lấy không thành công     | `Pickup failed`                    | Shipper không nhận được hàng từ người gửi (do lỗi hoặc vấn đề khác).          |
| Đã hủy giao                    | `Delivery canceled`                | Đơn hàng bị hủy trước khi giao cho người nhận.                                |
| VTP hủy lấy                   | `VTP pickup canceled`              | Đơn vị vận chuyển (VTP) hủy việc lấy hàng từ người gửi.                       |
| Shop hủy lấy                  | `Shop pickup canceled`             | Người gửi (shop) đã hủy việc giao hàng cho shipper.                          |
| Đang xác minh bồi thường       | `Under compensation review`        | Hàng hóa đang được xác minh để xử lý bồi thường.                              |
| Đã bồi thường                  | `Compensated`                      | Quá trình bồi thường hoàn tất, số tiền hoặc hàng hóa đã được xử lý.           |

---
