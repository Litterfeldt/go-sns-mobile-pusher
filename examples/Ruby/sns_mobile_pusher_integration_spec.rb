require 'spec_helper'

describe Integration::SNSMobilePusher do
  let(:success) { double(status: 200) }
  let(:failure) { double(status: 500) }
  let(:body) { { string: 'user_id', float: 132.2, int: 1 } }
  let(:string_body) { { string: 'user_id', float: '132.2', int: '1' } }

  describe '.push' do
    subject { described_class.push(body) }

    context 'when success' do
      before { mock_request(:post, 'send', string_body) { success } }

      it { is_expected.to be true }
    end

    context 'when failure' do
      before { mock_request(:post, 'send', string_body) { failure } }

      it { is_expected.to be false }
    end
  end

  def mock_request(method, url, body, &block)
    expect(Excon).to receive(method)
      .with(
        "/#{url}",
        body: JSON.dump(body),
        headers: {
          'Auth-Token' => ENV['PUSHER_TOKEN']
        },
        &block
      )
  end
end
