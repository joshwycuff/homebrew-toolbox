class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-5dc195e429bc157d295356a36e5c36712459e2f1/example-go.tar.gz"
  version "5dc195e429bc157d295356a36e5c36712459e2f1"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    bash_completion.install "example-go.bash" => "example-go"
    zsh_completion.install "example-go.zsh" => "_example-go"
    fish_completion.install "example-go.fish"
  end
end
