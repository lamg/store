require "test_helper"

class ProductTest < ActiveSupport::TestCase
  include ActionMailer::TestHelper

  test "sends email notifications when back in stock" do
    product = products(:tshirt)
    # set proudct out of stock
    product.update(inventory_count: 0)
    assert_emails 2 do
      product.update(inventory_count: 48)
    end
  end
end
